package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"github.com/dukfaar/dukfin/ent"
	"github.com/dukfaar/dukfin/ent/account"
	"github.com/dukfaar/dukfin/ent/currency"
	"github.com/dukfaar/dukfin/gui/dialogs"
	"github.com/dukfaar/dukfin/gui/views"
	"github.com/shopspring/decimal"

	_ "github.com/mattn/go-sqlite3"
)

func PopulateNewClient(client *ent.Client) {
	ctx := context.Background()
	client.Schema.Create(ctx)

	eur, err := client.Currency.Create().SetName("Euro").SetSymbol("EUR").Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Currency.Create().SetName("US-Dollar").SetSymbol("USD").Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	mainAcc, err := client.Account.Create().
		SetName("Main Account").
		SetCurrency(client.Currency.Query().Where(currency.SymbolEQ("EUR")).FirstX(ctx)).
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	funAcc, err := client.Account.Create().
		SetName("Fun Account").
		SetCurrency(client.Currency.Query().Where(currency.SymbolEQ("EUR")).FirstX(ctx)).
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Transaction.Create().
		SetCurrencyValue(decimal.NewFromInt(2900)).
		SetDate(time.Now().Add(time.Duration(time.Hour * -500))).
		SetToAccount(mainAcc).
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Transaction.Create().
		SetCurrencyValue(decimal.NewFromInt(50)).
		SetDate(time.Now().Add(time.Duration(time.Hour * -300))).
		SetFromAccount(mainAcc).
		SetToAccount(funAcc).
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	btc, err := client.Security.Create().SetName("Bitcoin").SetSymbol("BTC").SetCurrency(eur).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	client.SecurityPrice.Create().SetValue(decimal.NewFromInt(33)).SetSecurity(btc).SetDate(time.Now().Add(-time.Hour * 24 * 90)).Save(ctx)
	client.SecurityPrice.Create().SetValue(decimal.NewFromInt(30)).SetSecurity(btc).SetDate(time.Now().Add(-time.Hour * 24 * 30)).Save(ctx)
	client.SecurityPrice.Create().SetValue(decimal.NewFromInt(60)).SetSecurity(btc).SetDate(time.Now().Add(-time.Hour * 24 * 15)).Save(ctx)
	_, err = client.SecurityPrice.Create().SetValue(decimal.NewFromInt(300)).SetSecurity(btc).SetDate(time.Now()).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	mainPortfolio, err := client.Portfolio.Create().SetName("MainPortfolio").SetCurrency(eur).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Transaction.Create().
		SetSecurityAmount(decimal.NewFromInt(9)).
		SetCurrencyValue(decimal.NewFromInt(300)).
		SetDate(time.Now().Add(time.Duration(-time.Hour * 24 * 60))).
		SetFromAccount(mainAcc).
		SetToPortfolio(mainPortfolio).
		SetSecurity(btc).
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

const (
	LAST_OPENED_FILE_KEY = "lastOpenedFile"
)

func processRecurringTransaction(client *ent.Client, recurringTransaction *ent.RecurringTransaction) {
	for recurringTransaction.NextDate.Before(time.Now()) {
		fromAccount := recurringTransaction.QueryFromAccount().FirstX(context.Background())
		toAccount := recurringTransaction.QueryToAccount().FirstX(context.Background())
		fromPortfolio := recurringTransaction.QueryFromPortfolio().FirstX(context.Background())
		toPortfolio := recurringTransaction.QueryToPortfolio().FirstX(context.Background())
		security := recurringTransaction.QuerySecurity().FirstX(context.Background())
		create := client.Transaction.Create().
			SetCurrencyValue(recurringTransaction.CurrencyValue).
			SetDate(recurringTransaction.NextDate)
		if fromAccount != nil {
			create = create.SetFromAccount(fromAccount)
		}
		if toAccount != nil {
			create = create.SetToAccount(toAccount)
		}
		if fromPortfolio != nil {
			create = create.SetFromPortfolio(fromPortfolio)
		}
		if toPortfolio != nil {
			create = create.SetToPortfolio(toPortfolio)
		}
		if security != nil {
			create = create.SetSecurity(security)
		}
		create.Save(context.Background())
		updatedRecurringTransaction, err := recurringTransaction.Update().
			SetLastDate(recurringTransaction.NextDate).
			SetNextDate(recurringTransaction.NextDate.AddDate(recurringTransaction.IntervalYears, recurringTransaction.IntervalMonths, recurringTransaction.IntervalDays)).
			Save(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
		recurringTransaction = updatedRecurringTransaction
	}
}

func processRecurringTransactions(client *ent.Client) {
	recurringTransactions, err := client.RecurringTransaction.Query().All(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, recurringTransaction := range recurringTransactions {
		processRecurringTransaction(client, recurringTransaction)
	}
}

func openFile(app fyne.App, path string, win fyne.Window, newFile bool) {
	app.Preferences().SetString(LAST_OPENED_FILE_KEY, path)
	mode := "rw"
	if newFile {
		mode = "rwc"
	}
	client, err := ent.Open("sqlite3", "file:"+path+"?mode="+mode+"&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	win.SetContent(views.NewFileView(client))
	if newFile {
		PopulateNewClient(client)
	} else {
		client.Schema.Create(context.Background())
	}
	processRecurringTransactions(client)
}

func importAccountTransactions(r io.Reader, client *ent.Client) {
	csvReader := csv.NewReader(r)
	for line, err := csvReader.Read(); err == nil; line, err = csvReader.Read() {
		tx := client.Transaction.Create()
		if line[0] != "" {
			fromAccount, err := client.Account.Query().Where(account.Name(line[0])).First(context.Background())
			if err != nil {
				continue
			}
			tx.SetFromAccount(fromAccount)
		}
		if line[1] != "" {
			toAccount, err := client.Account.Query().Where(account.Name(line[1])).First(context.Background())
			if err != nil {
				continue
			}
			tx.SetToAccount(toAccount)
		}
		tx.SetCurrencyValue(decimal.RequireFromString(line[2]))
		date, err := time.Parse(time.RFC3339, line[3])
		if err != nil {
			continue
		}
		tx.SetDate(date)
		tx.Save(context.Background())
	}
}

var pprofFlag = flag.Bool("pprof", false, "Create a CPU-Profile")

func main() {
	flag.Parse()
	if *pprofFlag {
		fmt.Println("Profiling started")
		f, err := os.Create("cpuprof")
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
	}
	defer func() {
		if *pprofFlag {
			pprof.StopCPUProfile()
		}
	}()

	a := app.NewWithID("dukfaar.finance")
	w := a.NewWindow("DukFin")
	dialogs.SetMainWindow(w)

	extensionFileFilter := storage.NewExtensionFileFilter([]string{".df"})
	newFileDialog := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
		if uc == nil {
			return
		}
		uri := uc.URI()
		uc.Close()
		os.Remove(uri.Path())
		openPath := uri.Path()
		if !strings.HasSuffix(openPath, ".df") {
			openPath += ".df"
		}
		openFile(a, openPath, w, true)
	}, w)
	newFileDialog.SetFilter(extensionFileFilter)

	openFileDialog := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
		if uc == nil {
			return
		}
		uri := uc.URI()
		uc.Close()
		openFile(a, uri.Path(), w, false)
	}, w)
	openFileDialog.SetFilter(extensionFileFilter)

	importTransactionFileDialog := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
		if uc == nil {
			return
		}
		defer uc.Close()
		//importAccountTransactions(uc, client)
	}, w)
	importTransactionFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".csv"}))

	w.Resize(fyne.NewSize(800, 600))
	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New", func() { newFileDialog.Show() }),
			fyne.NewMenuItem("Open", func() { openFileDialog.Show() }),
		),
	))

	lastFile := a.Preferences().String(LAST_OPENED_FILE_KEY)
	if lastFile != "" {
		openFile(a, lastFile, w, false)
	}

	w.Show()
	a.Run()
}
