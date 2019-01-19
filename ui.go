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

	if event.Key() == tcell.KeyRune {
		if event.Rune() == 'q' {
			ui.application.Stop()
			return nil
		}
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

		jobOffersList:  NewOfferList(),
		jobOfferDetail: NewOfferView(),
	}

	ui.pagesWidget.AddPage("detail", ui.jobOfferDetail, true, true)
	ui.jobOfferDetail.SetOffer(&ui.context.offers[0])

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

// Run executes the graphical view for this application
func (ui *UserInterface) Run() error {
	ui.jobOffersList.SetOfferList(ui.context.offers)
	ui.application.SetInputCapture(ui.globalApplicationKeybidings)
	ui.application.SetRoot(ui.layout, true)
	ui.application.SetFocus(ui.pagesWidget)
	return ui.application.Run()
}
