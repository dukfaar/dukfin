package views

import (
	"context"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dukfaar/dukfin/ent"
	"github.com/dukfaar/dukfin/gui/renderer"
	"github.com/dukfaar/dukfin/gui/timeseries"
	"github.com/dukfaar/dukfin/gui/widgets"
	"github.com/shopspring/decimal"
)

type Projections struct {
	widget.BaseWidget

	client *ent.Client
}

func NewProjectionsView(client *ent.Client) *Projections {
	result := &Projections{
		client: client,
	}
	result.ExtendBaseWidget(result)
	return result
}

func (v *Projections) CreateRenderer() fyne.WidgetRenderer {
	accounts, err := v.client.Account.Query().WithCurrency().WithIncomingRecurringTransactions().WithOutgoingRecurringTransactions().All(context.Background())
	if err != nil {
		return nil
	}
	projectionTarget := time.Now().AddDate(30, 0, 0)
	timeSeries := widgets.NewTimeSeries()
	timeSeries.SetEndDate(projectionTarget)
	for i, account := range accounts {
		accountTs := timeseries.NewAccount(account)
		accountReccurringTs := timeseries.NewRecurringAccountTransactions(account.Edges.IncomingRecurringTransactions, account.Edges.OutgoingRecurringTransactions)
		accountStateTs := &timeseries.Sum{
			Series: []timeseries.TimeSeries{accountTs, accountReccurringTs},
		}
		interestTs := &timeseries.Interest{
			Series: accountStateTs,
			APR:    decimal.NewFromFloat(7.0),
		}
		newTs := &timeseries.Sum{
			Series: []timeseries.TimeSeries{accountStateTs, interestTs},
		}
		timeSeries.Series[account.Name] = renderer.Series{
			Color:  colorMap[i%len(colorMap)],
			Series: newTs,
		}
		newTs.Rebuild()
		newTs.GetValue(projectionTarget)
	}
	accountsTable := widget.NewTable(
		func() (int, int) { return len(accounts), 3 },
		func() fyne.CanvasObject { return widget.NewLabel("some long named stuff and things") },
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			account := accounts[tci.Row]
			switch tci.Col {
			case 0:
				co.(*widget.Label).SetText(account.Name)
			case 1:
				co.(*widget.Label).SetText(fmt.Sprintf("%f", timeSeries.Series[account.Name].Series.GetValue(projectionTarget).InexactFloat64()))
			case 2:
				co.(*widget.Label).SetText(account.Edges.Currency.Symbol)
			case 3:
				co.(*widget.Label).SetText(account.ExpectedAPR.String())
			}
		})

	renderer := &renderer.ContainerRenderer{
		Container: container.NewMax(container.NewVSplit(accountsTable, container.NewMax(timeSeries))),
	}
	return renderer
}
