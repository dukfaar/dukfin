package dialogs

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/dukfaar/dukfin/ent"
	"github.com/dukfaar/dukfin/gui/widgets"
)

type deposit struct {
	dialog.Dialog
	client *ent.Client

	transaction *ent.Transaction

	amountString  *widgets.DecimalEntry
	accountSelect *widgets.ItemSelect[*ent.Account]
	dateEntry     *widgets.TimeEntry
}

func (d *deposit) callbackCreate(confirmed bool) {
	if !confirmed {
		return
	}
	amount, err := d.amountString.GetDecimal()
	if err != nil {
		fmt.Println(err)
		return
	}
	date, err := d.dateEntry.GetTime()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = d.client.Transaction.Create().
		SetCurrencyValue(amount).
		SetDate(date).
		SetToAccount(d.accountSelect.Get()).
		Save(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (d *deposit) callbackUpdate(confirmed bool) {
	if !confirmed {
		return
	}
	amount, err := d.amountString.GetDecimal()
	if err != nil {
		fmt.Println(err)
		return
	}
	date, err := d.dateEntry.GetTime()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = d.transaction.Update().
		SetCurrencyValue(amount).
		SetDate(date).
		SetToAccount(d.accountSelect.Get()).
		Save(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
}

func NewDeposit(client *ent.Client) dialog.Dialog {
	result := deposit{
		client:        client,
		amountString:  widgets.NewDecimalEntry(),
		accountSelect: widgets.NewItemSelect(client.Account.Query().AllX(context.Background()), func(s *ent.Account) string { return s.Name }),
		dateEntry:     widgets.NewTimeEntry(),
	}
	items := []*widget.FormItem{
		widget.NewFormItem("Account", result.accountSelect),
		widget.NewFormItem("Amount", result.amountString),
		widget.NewFormItem("Date", result.dateEntry),
	}
	return dialog.NewForm("Deposit", "Create", "Cancel", items, result.callbackCreate, window)
}

func EditDeposit(client *ent.Client, tx *ent.Transaction) dialog.Dialog {
	result := deposit{
		client:        client,
		transaction:   tx,
		amountString:  widgets.NewDecimalEntry(),
		accountSelect: widgets.NewItemSelect(client.Account.Query().AllX(context.Background()), func(s *ent.Account) string { return s.Name }),
		dateEntry:     widgets.NewTimeEntry(),
	}
	result.amountString.SetDecimal(tx.CurrencyValue)
	result.accountSelect.Set(tx.QueryToAccount().FirstX(context.Background()))
	result.dateEntry.SetTime(tx.Date)
	items := []*widget.FormItem{
		widget.NewFormItem("Account", result.accountSelect),
		widget.NewFormItem("Amount", result.amountString),
		widget.NewFormItem("Date", result.dateEntry),
	}
	return dialog.NewForm("Deposit", "Update", "Cancel", items, result.callbackUpdate, window)
}
