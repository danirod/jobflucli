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

	// Wrap the title and status labels to make them full width.
	titleFrame  *tview.Frame
	statusFrame *tview.Frame

	// Page widgets
	jobOffersList *OfferList

	// Flex layout
	layout *tview.Flex
}

func (ui *UserInterface) applyTheme() {
	ui.titleFrame.SetBackgroundColor(tcell.ColorBlue)
	ui.titleFrame.SetBorder(false)
	ui.titleWidget.SetBackgroundColor(tcell.ColorDefault)
	ui.titleWidget.SetTextColor(tcell.ColorYellow)

	ui.statusFrame.SetBackgroundColor(tcell.ColorGreen)
	ui.statusFrame.SetBorder(false)
	ui.statusWidget.SetBackgroundColor(tcell.ColorDefault)
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
		titleFrame:   tview.NewFrame(titleLabel),
		statusFrame:  tview.NewFrame(statusLabel),

		pagesWidget: tview.NewPages(),

		layout: tview.NewFlex(),
	}

	ui.jobOffersList = NewOfferList()

	ui.pagesWidget.AddPage("jobs", ui.jobOffersList, true, true)

	ui.layout.SetDirection(tview.FlexRow)
	ui.layout.AddItem(ui.titleWidget, 1, 1, false)
	ui.layout.AddItem(ui.pagesWidget, 0, 1, false)
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
