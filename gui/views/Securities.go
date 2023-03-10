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
)

type Securities struct {
	widget.BaseWidget

	client *ent.Client
}

func NewSecuritiesView(client *ent.Client) *Securities {
	result := &Securities{
		client: client,
	}
	result.ExtendBaseWidget(result)
	return result
}

func (v *Securities) CreateRenderer() fyne.WidgetRenderer {
	Securities, err := v.client.Security.Query().WithCurrency().All(context.Background())
	if err != nil {
		return nil
	}
	timeSeries := widgets.NewTimeSeries()
	for i, security := range Securities {
		newTs := timeseries.NewSecurityPrice(security)
		timeSeries.Series[security.Name] = renderer.Series{
			Color:  colorMap[i%len(colorMap)],
			Series: newTs,
		}
		newTs.Rebuild()
	}
	SecuritiesTable := widget.NewTable(
		func() (int, int) { return len(Securities), 3 },
		func() fyne.CanvasObject { return widget.NewLabel("some long named stuff and things") },
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			security := Securities[tci.Row]
			switch tci.Col {
			case 0:
				co.(*widget.Label).SetText(security.Name)
			case 1:
				co.(*widget.Label).SetText(fmt.Sprintf("%f", timeSeries.Series[security.Name].Series.GetValue(time.Now()).InexactFloat64()))
			case 2:
				co.(*widget.Label).SetText(security.Edges.Currency.Symbol)
			}
		})

	renderer := &renderer.ContainerRenderer{
		Container: container.NewMax(container.NewVSplit(SecuritiesTable, container.NewMax(timeSeries))),
	}
	return renderer
}
