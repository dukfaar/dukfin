package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dukfaar/dukfin/ent"
	"github.com/dukfaar/dukfin/gui/renderer"
)

type File struct {
	widget.BaseWidget

	client *ent.Client
}

func NewFileView(client *ent.Client) *File {
	result := &File{
		client: client,
	}
	result.ExtendBaseWidget(result)
	return result
}

func (v *File) CreateRenderer() fyne.WidgetRenderer {
	accountsButton := widget.NewButton("Accounts", nil)
	portfoliosButton := widget.NewButton("Portfolios", nil)
	recurringTransactions := widget.NewButton("Recurring Transactions", nil)
	stocksButton := widget.NewButton("Stocks", nil)
	stockPositionsButton := widget.NewButton("StockPositions", nil)
	currenciesButton := widget.NewButton("Currencies", nil)
	projectionsButton := widget.NewButton("Projections", nil)

	mainView := widget.NewLabel("MainView")
	sideMenu := container.NewVBox(
		accountsButton,
		portfoliosButton,
		recurringTransactions,
		stocksButton,
		stockPositionsButton,
		currenciesButton,
		projectionsButton,
	)
	renderer := &renderer.ContainerRenderer{
		Container: container.NewBorder(nil, nil, sideMenu, nil, mainView),
	}

	accountsButton.OnTapped = func() {
		renderer.Container.Objects[0] = NewAccountsView(v.client)
		renderer.Refresh()
	}
	portfoliosButton.OnTapped = func() {
		renderer.Container.Objects[0] = NewPortfoliosView(v.client)
		renderer.Refresh()
	}
	stocksButton.OnTapped = func() {
		renderer.Container.Objects[0] = NewSecuritiesView(v.client)
		renderer.Refresh()
	}
	projectionsButton.OnTapped = func() {
		renderer.Container.Objects[0] = NewProjectionsView(v.client)
		renderer.Refresh()
	}
	recurringTransactions.OnTapped = func() {
		renderer.Container.Objects[0] = NewRecurringTransactionsView(v.client)
		renderer.Refresh()
	}
	return renderer
}
