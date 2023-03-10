package dialogs

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/dukfaar/dukfin/ent"
	"github.com/dukfaar/dukfin/gui/widgets"
)

func NewDeposit(client *ent.Client) dialog.Dialog {
	amountString := widgets.NewDecimalEntry()
	accountSelect := widgets.NewItemSelect(client.Account.Query().AllX(context.Background()), func(s *ent.Account) string { return s.Name })
	dateEntry := widgets.NewTimeEntry()
	items := []*widget.FormItem{
		widget.NewFormItem("Account", accountSelect),
		widget.NewFormItem("Amount", amountString),
		widget.NewFormItem("Date", dateEntry),
	}
	callback := func(confirmed bool) {
		if !confirmed {
			return
		}
		amount, err := amountString.GetDecimal()
		if err != nil {
			fmt.Println(err)
			return
		}
		date, err := dateEntry.GetTime()
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = client.Transaction.Create().
			SetCurrencyValue(amount).
			SetDate(date).
			SetToAccount(accountSelect.Get()).
			Save(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	return dialog.NewForm("Deposit", "Create", "Cancel", items, callback, window)
}
