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
)

type RecurringTransactions struct {
	widget.BaseWidget

	client *ent.Client
}

func NewRecurringTransactionsView(client *ent.Client) *RecurringTransactions {
	result := &RecurringTransactions{
		client: client,
	}
	result.ExtendBaseWidget(result)
	return result
}

func (v *RecurringTransactions) CreateRenderer() fyne.WidgetRenderer {
	recurringTransactions, err := v.client.RecurringTransaction.
		Query().
		WithFromAccount().
		WithToAccount().
		WithFromPortfolio().
		WithToPortfolio().
		WithSecurity().
		All(context.Background())
	if err != nil {
		return nil
	}
	accountToolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.MoveDownIcon(), func() { dialogs.NewRecurringDeposit(v.client).Show() }),
		widget.NewToolbarAction(theme.MoveUpIcon(), func() { dialogs.NewRecurringWithdrawal(v.client).Show() }),
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() { dialogs.NewRecurringTransfer(v.client).Show() }),
	)
	RecurringTransactionsTable := widget.NewTable(
		func() (int, int) { return len(recurringTransactions), 4 },
		func() fyne.CanvasObject { return widget.NewLabel("some long named stuff and things") },
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			recurringTransaction := recurringTransactions[tci.Row]
			switch tci.Col {
			case 0:
				co.(*widget.Label).SetText(recurringTransaction.Name)
			case 1:
				co.(*widget.Label).SetText(recurringTransaction.CurrencyValue.String())
			case 2:
				if recurringTransaction.LastDate != nil {
					co.(*widget.Label).SetText(recurringTransaction.LastDate.Format(time.RFC3339))
				}
			case 3:
				co.(*widget.Label).SetText(recurringTransaction.NextDate.Format(time.RFC3339))
			case 4:
				co.(*widget.Label).SetText(fmt.Sprintf("%d years, %d months, %d days", recurringTransaction.IntervalYears, recurringTransaction.IntervalMonths, recurringTransaction.IntervalDays))
			case 5:
				if recurringTransaction.Edges.FromAccount != nil {
					co.(*widget.Label).SetText(recurringTransaction.Edges.FromAccount.Name)
				} else if recurringTransaction.Edges.FromPortfolio != nil {
					co.(*widget.Label).SetText(recurringTransaction.Edges.FromPortfolio.Name)
				}
			case 6:
				if recurringTransaction.Edges.ToAccount != nil {
					co.(*widget.Label).SetText(recurringTransaction.Edges.ToAccount.Name)
				} else if recurringTransaction.Edges.ToPortfolio != nil {
					co.(*widget.Label).SetText(recurringTransaction.Edges.ToPortfolio.Name)
				}
			case 7:
				if recurringTransaction.Edges.Security != nil {
					co.(*widget.Label).SetText(recurringTransaction.Edges.Security.Name)
				}
			}
		})
	renderer := &renderer.ContainerRenderer{
		Container: container.NewMax(
			container.NewBorder(
				accountToolbar, nil, nil, nil,
				RecurringTransactionsTable,
			),
		),
	}
	return renderer
}
