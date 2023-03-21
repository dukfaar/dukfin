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
	"github.com/dukfaar/dukfin/gui/widgets"
)

type Accounts struct {
	widget.BaseWidget

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
	timeSeries := widgets.NewTimeSeries()
	for _, account := range accounts {
		newTs := timeseries.NewAccount(account)
		timeSeries.Series[account.Name] = newTs
		newTs.Rebuild()
	}
	accountToolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() { dialogs.Account(v.client, nil).Show() }),
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
				co.(*widget.Label).SetText(fmt.Sprintf("%f", timeSeries.Series[account.Name].GetValue(time.Now()).InexactFloat64()))
			case 2:
				co.(*widget.Label).SetText(account.Edges.Currency.Symbol)
			case 3:
				co.(*widget.Label).SetText(account.ExpectedAPR.String())
			}
		})
	accountsTable.OnSelected = func(tci widget.TableCellID) {
		dialogs.Account(v.client, accounts[tci.Row]).Show()
	}
	renderer := &renderer.ContainerRenderer{
		Container: container.NewMax(container.NewVSplit(
			container.NewBorder(
				accountToolbar, nil, nil, nil,
				accountsTable,
			),
			container.NewMax(timeSeries),
		)),
	}
	return renderer
}
