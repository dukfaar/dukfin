package dialogs

import (
	"context"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/dukfaar/dukfin/ent"
	"github.com/dukfaar/dukfin/gui/widgets"
)

func NewPortfolio(client *ent.Client) dialog.Dialog {
	name := widget.NewEntry()
	currencySelect := widgets.NewItemSelect(client.Currency.Query().AllX(context.Background()), func(c *ent.Currency) string { return c.Name })
	referenceAccountSelect := widgets.NewItemSelect(client.Account.Query().AllX(context.Background()), func(a *ent.Account) string { return a.Name })
	items := []*widget.FormItem{
		widget.NewFormItem("Name", name),
		widget.NewFormItem("Currency", currencySelect),
		widget.NewFormItem("Reference Account", referenceAccountSelect),
	}
	callback := func(confirmed bool) {
		if !confirmed {
			return
		}
		client.Portfolio.Create().
			SetCurrency(currencySelect.Get()).
			SetName(name.Text).
			SetReferenceAccount(referenceAccountSelect.Get()).
			Save(context.Background())
	}
	return dialog.NewForm("Portfolio", "Create", "Cancel", items, callback, window)
}
