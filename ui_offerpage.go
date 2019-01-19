package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"strings"
)

// OfferView is a tview widget used to represent an offer page. It paints the
// information about the widget on the top of the page, and the remaining
// available space is filled with a text area with the contents of the offer.
type OfferView struct {
	*tview.Grid
	offer             *Offer           // offer being displayed
	positionWidget    *tview.TableCell // will contain the offer position
	companyWidget     *tview.TableCell // will contain the company field
	dateWidget        *tview.TableCell // date at which the offer was posted
	tagsWidget        *tview.TableCell // tags in the offer
	urlWidget         *tview.TableCell // link to open the offer in a browser
	descriptionWidget *tview.TextView  // main content of the offer
}

// SetOffer will render the given offer in the widget. This function should be
// called before presenting the widget to the user so that the data can be loaded.
// This function updates the widgets, so to avoid race conditions it should be
// called in the main thread or using a repaint event.
func (ov *OfferView) SetOffer(offer *Offer) {
	ov.offer = offer
	ov.positionWidget.SetText(strings.TrimSpace(offer.Position))
	ov.companyWidget.SetText(strings.TrimSpace(offer.Company))
	ov.dateWidget.SetText(offer.CreationDate.Format("Mon, 2 Jan 2006 15:04:05"))
	ov.tagsWidget.SetText(strings.Join(offer.Tags, ", "))
	ov.urlWidget.SetText(offer.URL)
	ov.descriptionWidget.SetText(offer.Description)
}

// NewOfferView initialises a new widget to be used as a page to present
// information about an offer. The widget will have a header section with
// information about the offer itself, and a main area section with the
// description of the offer.
func NewOfferView() *OfferView {
	offerView := &OfferView{
		Grid:           tview.NewGrid(),
		positionWidget: tview.NewTableCell("").SetExpansion(1),
		companyWidget:  tview.NewTableCell("").SetExpansion(1),
		dateWidget:     tview.NewTableCell("").SetExpansion(1),
		tagsWidget:     tview.NewTableCell("").SetExpansion(1),
		urlWidget:      tview.NewTableCell("").SetExpansion(1),
	}

	// The header table has information about the offer.
	headerTable := NewHeaderTable()
	headerTable.AddRow("Position:", offerView.positionWidget)
	headerTable.AddRow("Company:", offerView.companyWidget)
	headerTable.AddRow("Date:", offerView.dateWidget)
	headerTable.AddRow("Tags:", offerView.tagsWidget)
	headerTable.AddRow("URL:", offerView.urlWidget)

	// The description widget renders the offer content.
	descriptionWidget := tview.NewTextView().SetWordWrap(true).SetScrollable(true)
	descriptionSizingFunc := func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		ix, iy, iw, ih := offerView.descriptionWidget.GetInnerRect()
		if iw > 65 {
			return ix, iy, 65, ih
		}
		return ix, iy, iw, ih
	}
	descriptionWidget.SetDrawFunc(descriptionSizingFunc)
	offerView.descriptionWidget = descriptionWidget

	// Reflow the stuff.
	offerView.SetRows(5, 1, -1)
	offerView.AddItem(headerTable, 0, 0, 1, 1, 0, 0, false)
	offerView.AddItem(tview.NewBox().SetBackgroundColor(tcell.ColorSilver), 1, 0, 1, 1, 0, 0, false)
	offerView.AddItem(offerView.descriptionWidget, 2, 0, 1, 1, 0, 0, true)
	return offerView
}
