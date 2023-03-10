package dialogs

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/dukfaar/dukfin/ent"
	"github.com/dukfaar/dukfin/gui/widgets"
)

func NewRecurringTransfer(client *ent.Client) dialog.Dialog {
	nameEntry := widget.NewEntry()
	amountString := widgets.NewDecimalEntry()
	accounts := client.Account.Query().AllX(context.Background())
	toAccountSelect := widgets.NewItemSelect(accounts, func(s *ent.Account) string { return s.Name })
	fromAccountSelect := widgets.NewItemSelect(accounts, func(s *ent.Account) string { return s.Name })
	dateEntry := widgets.NewTimeEntry()
	daysEntry := widgets.NewIntEntry()
	daysEntry.SetText("0")
	monthsEntry := widgets.NewIntEntry()
	monthsEntry.SetText("1")
	yearsEntry := widgets.NewIntEntry()
	yearsEntry.SetText("0")
	items := []*widget.FormItem{
		widget.NewFormItem("Name", nameEntry),
		widget.NewFormItem("From", fromAccountSelect),
		widget.NewFormItem("To", toAccountSelect),
		widget.NewFormItem("Amount", amountString),
		widget.NewFormItem("Next Date", dateEntry),
		widget.NewFormItem("Every x days", daysEntry),
		widget.NewFormItem("Every x months", monthsEntry),
		widget.NewFormItem("Every x years", yearsEntry),
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
		days, err := daysEntry.GetInt()
		if err != nil {
			fmt.Println(err)
			return
		}
		months, err := monthsEntry.GetInt()
		if err != nil {
			fmt.Println(err)
			return
		}
		years, err := yearsEntry.GetInt()
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = client.RecurringTransaction.Create().
			SetName(nameEntry.Text).
			SetCurrencyValue(amount).
			SetNextDate(date).
			SetFromAccount(fromAccountSelect.Get()).
			SetToAccount(toAccountSelect.Get()).
			SetIntervalDays(days).
			SetIntervalMonths(months).
			SetIntervalYears(years).
			Save(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	return dialog.NewForm("Recurring Deposit", "Create", "Cancel", items, callback, window)
}
