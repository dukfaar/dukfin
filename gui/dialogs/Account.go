package dialogs

import (
	"context"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/dukfaar/dukfin/ent"
	"github.com/dukfaar/dukfin/gui/widgets"
)

func Account(client *ent.Client, account *ent.Account) dialog.Dialog {
	accountName := widget.NewEntry()
	expectedAprString := widgets.NewDecimalEntry()
	currencySelect := widgets.NewItemSelect(client.Currency.Query().AllX(context.Background()), func(s *ent.Currency) string { return s.Name })
	if account != nil {
		accountName.SetText(account.Name)
		expectedAprString.SetText(account.ExpectedAPR.String())
		currencySelect.Set(account.QueryCurrency().FirstX(context.Background()))
	}
	items := []*widget.FormItem{
		widget.NewFormItem("Name", accountName),
		widget.NewFormItem("Currency", currencySelect),
		widget.NewFormItem("Expected APR", expectedAprString),
	}
	callback := func(confirmed bool) {
		if !confirmed {
			return
		}
		apr, err := expectedAprString.GetDecimal()
		if err != nil {
			return
		}
		if account != nil {
			account.Update().
				SetCurrency(currencySelect.Get()).
				SetName(accountName.Text).
				SetExpectedAPR(apr).
				Save(context.Background())
		} else {
			client.Account.Create().
				SetCurrency(currencySelect.Get()).
				SetName(accountName.Text).
				SetExpectedAPR(apr).
				Save(context.Background())
		}
	}
	return dialog.NewForm("Account", "Create", "Cancel", items, callback, window)
}
