package main

import (
	"fmt"
	"strings"

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
	jobList           *tview.Table
	jobListBackingIds map[int]int

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

func newJobListWidget(ui *UserInterface) *tview.Table {
	table := tview.NewTable()
	table.SetSelectable(true, false)
	table.SetSelectedFunc(func(row, col int) {
		backOffer := ui.jobListBackingIds[row]
		ui.onOfferSelected(backOffer)
	})
	return table
}

func (ui *UserInterface) onOfferSelected(offerID int) {
	panic(fmt.Errorf("Selected %v", ui.context.GetOffer(offerID)))
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

	ui.jobList = newJobListWidget(ui)

	ui.pagesWidget.AddPage("jobs", ui.jobList, true, true)

	ui.layout.SetDirection(tview.FlexRow)
	ui.layout.AddItem(ui.titleWidget, 1, 1, false)
	ui.layout.AddItem(ui.pagesWidget, 0, 1, false)
	ui.layout.AddItem(ui.statusWidget, 1, 1, false)

	ui.applyTheme()

	return ui
}

func (ui *UserInterface) reloadJobOffers() {
	ui.jobList.Clear()
	ui.jobListBackingIds = make(map[int]int)

	for i, offer := range ui.context.offers {
		// Format the offer to make it look nice.
		position := strings.TrimSpace(offer.Position)
		company := strings.TrimSpace(offer.Company)
		offerCell := tview.NewTableCell(fmt.Sprintf("%s at %s", position, company))
		offerCell.SetExpansion(1)

		// Then add the offer to the table.
		ui.jobList.SetCellSimple(i, 0, offer.CreationDate.String())
		ui.jobList.SetCellSimple(i, 1, " ")
		ui.jobList.SetCell(i, 2, offerCell)

		// Put a backing ID so that we can refer to this offer later.
		ui.jobListBackingIds[i] = offer.ID
	}
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
	ui.reloadJobOffers()
	ui.application.SetInputCapture(ui.globalApplicationKeybidings)
	ui.application.SetRoot(ui.layout, true)
	ui.application.SetFocus(ui.pagesWidget)
	return ui.application.Run()
}
