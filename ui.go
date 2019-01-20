package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// UserInterface represents the TUI used by this application.
type UserInterface struct {
	application  *tview.Application
	context      *Context
	titleWidget  *tview.TextView
	pagesWidget  *tview.Pages
	statusWidget *tview.TextView

	// Page widgets
	locationsList  *LocationsTable
	jobOffersList  *OfferList
	jobOfferDetail *OfferView

	// Flex layout
	layout *tview.Flex
}

func (ui *UserInterface) applyTheme() {
	ui.titleWidget.SetBackgroundColor(tcell.ColorBlue)
	ui.titleWidget.SetTextColor(tcell.ColorYellow)
	ui.statusWidget.SetBackgroundColor(tcell.ColorGreen)
	ui.statusWidget.SetTextColor(tcell.ColorYellow)
	ui.jobOffersList.SetSelectedStyle(tcell.ColorWhite, tcell.ColorBlue, tcell.AttrNone)
}

func (ui *UserInterface) globalApplicationKeybidings(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyCtrlL {
		ui.application.Draw()
		return nil
	}
	return event
}

// NewUserInterface creates a new user interface given a context state.
func NewUserInterface(context *Context) *UserInterface {
	titleLabel := tview.NewTextView()
	statusLabel := tview.NewTextView()

	ui := &UserInterface{
		application: tview.NewApplication(),
		context:     context,

		titleWidget:  titleLabel,
		statusWidget: statusLabel,

		pagesWidget: tview.NewPages(),

		layout: tview.NewFlex(),

		locationsList:  NewLocationsTable(),
		jobOffersList:  NewOfferList(),
		jobOfferDetail: NewOfferView(),
	}

	ui.jobOffersList.SetSelectedFunc(func(row, col int) {
		// Get the selected offer by looking the reverse map.
		offerID := ui.jobOffersList.backingOfferIds[row]
		offer := ui.context.GetOffer(offerID)
		ui.SwitchToOffer(offer)
	})

	ui.locationsList.SetSelectedFunc(func(row, col int) {
		// Get the location and fetch offers for that location.
		location := ui.locationsList.GetSelectedLocation()
		if err := ui.context.SetOffersByLocation(location); err != nil {
			panic(err)
		}
		ui.SwitchToList()
	})

	ui.locationsList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune && event.Rune() == 'q' {
			ui.application.Stop()
			return event
		}
		return event
	})

	ui.jobOffersList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune && event.Rune() == 'q' {
			ui.application.Stop()
			return event
		}
		if event.Key() == tcell.KeyRune && event.Rune() == 'l' {
			ui.SwitchToLocations()
			return event
		}
		return event
	})

	ui.jobOfferDetail.descriptionWidget.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune && event.Rune() == 'q' {
			ui.SwitchToList()
		}
		return event
	})

	ui.pagesWidget.AddPage("locations", ui.locationsList, true, false)
	ui.pagesWidget.AddPage("list", ui.jobOffersList, true, false)
	ui.pagesWidget.AddPage("detail", ui.jobOfferDetail, true, false)
	ui.pagesWidget.ShowPage("list")
	ui.application.SetFocus(ui.jobOffersList)

	ui.layout.SetDirection(tview.FlexRow)
	ui.layout.AddItem(ui.titleWidget, 1, 1, false)
	ui.layout.AddItem(ui.pagesWidget, 0, 1, true)
	ui.layout.AddItem(ui.statusWidget, 1, 1, false)

	ui.applyTheme()

	return ui
}

func (ui *UserInterface) SetTitle(title string) {
	ui.application.QueueUpdateDraw(func() {
		ui.titleWidget.Clear()
		ui.titleWidget.SetText(title)
	})
}

func (ui *UserInterface) SetStatus(status string) {
	ui.application.QueueUpdateDraw(func() {
		ui.statusWidget.Clear()
		ui.statusWidget.SetText(status)
	})
}

func (ui *UserInterface) SwitchToLocations() {
	ui.pagesWidget.SwitchToPage("locations")
	ui.application.SetFocus(ui.locationsList)
	ui.SetTitle("JobFluCli | Select a location")
	ui.SetStatus("q:Quit   j/Up:MoveUp   k/Down:MoveDown")
}

func (ui *UserInterface) SwitchToList() {
	ui.pagesWidget.SwitchToPage("list")
	ui.jobOffersList.SetOfferList(ui.context.offers)
	ui.application.SetFocus(ui.jobOffersList)
	ui.SetTitle("JobFluCli | List of offers")
	ui.SetStatus("q:Quit   j/Up:MoveUp   k/Down:MoveDown   l:SwitchLocation")
}

func (ui *UserInterface) SwitchToOffer(o *Offer) {
	ui.jobOfferDetail.SetOffer(o)
	ui.pagesWidget.SwitchToPage("detail")
	ui.application.SetFocus(ui.jobOfferDetail)
	ui.SetTitle("JobFluCli | Offer Information")
	ui.SetStatus("q:Back   j/Up:MoveUp   k/Down:MoveDown")
}

// Run executes the graphical view for this application
func (ui *UserInterface) Run() error {
	ui.application.SetInputCapture(ui.globalApplicationKeybidings)
	ui.application.SetRoot(ui.layout, true)
	ui.application.SetFocus(ui.pagesWidget)
	return ui.application.Run()
}
