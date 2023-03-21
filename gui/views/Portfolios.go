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

type Portfolios struct {
	widget.BaseWidget

	client *ent.Client
}

func NewPortfoliosView(client *ent.Client) *Portfolios {
	result := &Portfolios{
		client: client,
	}
	result.ExtendBaseWidget(result)
	return result
}

func (v *Portfolios) CreateRenderer() fyne.WidgetRenderer {
	portfolios, err := v.client.Portfolio.Query().WithCurrency().All(context.Background())
	if err != nil {
		return nil
	}
	securities, err := v.client.Security.Query().All(context.Background())
	if err != nil {
		return nil
	}
	timeSeries := widgets.NewTimeSeries()
	for _, portfolio := range portfolios {
		ts := &timeseries.Sum{Series: []timeseries.TimeSeries{}}
		for _, security := range securities {
			newTs := &timeseries.Multiply{
				Series1: timeseries.NewPortfolioSecurityAmount(portfolio, security),
				Series2: timeseries.NewSecurityPrice(security),
			}
			ts.Series = append(ts.Series, newTs)
		}
		ts.Rebuild()
		timeSeries.Series[portfolio.Name] = ts
	}
	portfolioToolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() { dialogs.NewPortfolio(v.client).Show() }),
	)
	portfolioTable := widget.NewTable(
		func() (int, int) { return len(portfolios), 4 },
		func() fyne.CanvasObject { return widget.NewLabel("some long named stuff and things") },
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			portfolio := portfolios[tci.Row]
			switch tci.Col {
			case 0:
				co.(*widget.Label).SetText(portfolio.Name)
			case 1:
				co.(*widget.Label).SetText(fmt.Sprintf("%f", timeSeries.Series[portfolio.Name].GetValue(time.Now()).InexactFloat64()))
			case 2:
				co.(*widget.Label).SetText(portfolio.Edges.Currency.Symbol)
			case 3:
				if portfolio.Edges.ReferenceAccount == nil {
					co.(*widget.Label).SetText("")
					return
				}
				co.(*widget.Label).SetText(portfolio.Edges.ReferenceAccount.Name)
			}
		},
	)
	renderer := &renderer.ContainerRenderer{
		Container: container.NewMax(container.NewVSplit(
			container.NewBorder(
				portfolioToolbar, nil, nil, nil,
				portfolioTable,
			),
			container.NewMax(timeSeries),
		)),
	}
	return renderer
}
