package views

import (
	"context"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dukfaar/dukfin/ent"
	"github.com/dukfaar/dukfin/gui/dialogs"
	"github.com/dukfaar/dukfin/gui/renderer"
	"github.com/dukfaar/dukfin/gui/timeseries"
)

type Accounts struct {
	widget.BaseWidget

	//transactionTable *widget.Table

	client *ent.Client
}

func NewAccountsView(client *ent.Client) *Accounts {
	result := &Accounts{
		client: client,
	}
	result.ExtendBaseWidget(result)
	return result
}

func (v *Accounts) CreateRenderer() fyne.WidgetRenderer {
	accounts, err := v.client.Account.Query().WithCurrency().All(context.Background())
	if err != nil {
		return nil
	}
	timeseriesMap := make(map[int]timeseries.TimeSeries)
	for _, account := range accounts {
		newTs := timeseries.NewAccount(account)
		timeseriesMap[account.ID] = newTs
		newTs.Rebuild()
	}
	accountToolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() { dialogs.NewAccount(v.client, nil).Show() }),
		widget.NewToolbarAction(theme.MoveDownIcon(), func() { dialogs.NewDeposit(v.client).Show() }),
		widget.NewToolbarAction(theme.MoveUpIcon(), func() { dialogs.NewWithdrawal(v.client).Show() }),
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() { dialogs.NewTransfer(v.client).Show() }),
	)
	accountsTable := widget.NewTable(
		func() (int, int) { return len(accounts), 4 },
		func() fyne.CanvasObject { return widget.NewLabel("some long named stuff and things") },
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			account := accounts[tci.Row]
			switch tci.Col {
			case 0:
				co.(*widget.Label).SetText(account.Name)
			case 1:
				co.(*widget.Label).SetText(fmt.Sprintf("%f", timeseriesMap[account.ID].GetValue(time.Now()).InexactFloat64()))
			case 2:
				co.(*widget.Label).SetText(account.Edges.Currency.Symbol)
			case 3:
				co.(*widget.Label).SetText(account.ExpectedAPR.String())
			}
		})
	accountBorder := container.NewBorder(
		accountToolbar, nil, nil, nil,
		accountsTable,
	)
	maxContainer := container.NewMax(accountBorder)
	widgetRenderer := &renderer.ContainerRenderer{Container: maxContainer}
	accountsTable.OnSelected = func(tci widget.TableCellID) {
		account := accounts[tci.Row]
		txsAll := timeseriesMap[account.ID].(*timeseries.Account).GetAllAccountTransactionsInOrder(time.Now())
		maxContainer.Objects = []fyne.CanvasObject{container.NewVSplit(
			accountBorder,
			container.NewMax(widget.NewTable(
				func() (int, int) { return len(txsAll), 6 },
				func() fyne.CanvasObject { return widget.NewLabel("some long named stuff and things") },
				func(tci widget.TableCellID, co fyne.CanvasObject) {
					ctx := context.Background()
					tx := txsAll[tci.Row]
					fromAccount := tx.QueryFromAccount().FirstX(ctx)
					toAccount := tx.QueryToAccount().FirstX(ctx)
					fromPortfolio := tx.QueryFromPortfolio().FirstX(ctx)
					toPortfolio := tx.QueryToPortfolio().FirstX(ctx)
					switch tci.Col {
					case 0:
						co.(*widget.Label).SetText(tx.CurrencyValue.StringFixed(2))
					case 1:
						if toAccount != nil && toAccount.ID == account.ID {
							if fromPortfolio != nil {
								co.(*widget.Label).SetText("Sell")
							} else if fromAccount != nil {
								co.(*widget.Label).SetText("Transfer In")
							} else {
								co.(*widget.Label).SetText("Deposit")
							}
						} else if fromAccount != nil && fromAccount.ID == account.ID {
							if toPortfolio != nil {
								co.(*widget.Label).SetText("Buy")
							} else if toAccount != nil {
								co.(*widget.Label).SetText("Transfer Out")
							} else {
								co.(*widget.Label).SetText("Withdrawal")
							}
						}
					case 2:
						co.(*widget.Label).SetText(fmt.Sprintf("%f", timeseriesMap[account.ID].GetValue(tx.Date).InexactFloat64()))
					case 3:
						if toAccount != nil && toAccount.ID == account.ID {
							if fromPortfolio != nil {
								co.(*widget.Label).SetText(fromPortfolio.Name)
								return
							}
							if fromAccount != nil {
								co.(*widget.Label).SetText(fromAccount.Name)
								return
							}
						} else if fromAccount != nil && fromAccount.ID == account.ID {
							if toPortfolio != nil {
								co.(*widget.Label).SetText(toPortfolio.Name)
								return
							}
							if toAccount != nil {
								co.(*widget.Label).SetText(toAccount.Name)
								return
							}
						}
						co.(*widget.Label).SetText("")
					case 4:
						if security := tx.QuerySecurity().FirstX(ctx); security != nil {
							co.(*widget.Label).SetText(security.Name)
						} else {
							co.(*widget.Label).SetText("")
						}
					case 5:
						co.(*widget.Label).SetText(tx.Date.Format(time.RFC3339))
					}
				})),
		)}
	}
	return widgetRenderer
}
