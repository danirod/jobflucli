package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"strings"
)

// OfferList is a widget that represents the list of offers downloaded.
type OfferList struct {
	*tview.Table
	backingOffers   []Offer                // the offers themselves.
	backingOfferIds map[int]int            // maps each row with the underlying offer
	filterFunc      func(offer Offer) bool // used to filter the presented offers
}

// NewOfferList returns a new offerlist widget that can be used to present offers.
func NewOfferList() *OfferList {
	offerList := new(OfferList)
	offerList.Table = tview.NewTable()
	offerList.SetSelectable(true, false)
	offerList.filterFunc = nil
	return offerList
}

// SetFilterFunc allows to filter the offers displayed in the table.
// This function modifies the widget. Therefore, it is required to either
// call this in the mainthread, or dispatch an event to the application.
func (ol *OfferList) SetFilterFunc(filter func(offer Offer) bool) {
	ol.filterFunc = filter
	ol.reloadTable()
}

// SetOfferList updates the list of offers presented to the users.
// This function modifies the widget. Therefore, it is required to either
// call this in the mainthread, or dispatch an event to the application.
func (ol *OfferList) SetOfferList(offers []Offer) {
	ol.backingOffers = offers
	ol.reloadTable()
}

// SetOfferList will update the table contained in the offer list by the list
// of offers given as an argument. It will also update every handler so that
// the new items of the table can be selected to toggle them.
func (ol *OfferList) reloadTable() {
	// Clear the selection model.
	ol.Clear()
	ol.backingOfferIds = make(map[int]int)

	// Put the new selection model.
	nextRow := 0
	for _, offer := range ol.backingOffers {
		if ol.filterFunc != nil && !ol.filterFunc(offer) {
			continue
		}

		// Format timestamp
		timestamp := offer.CreationDate.Format("2006 Jan 2, 15:04")
		timestampCell := tview.NewTableCell(timestamp)
		timestampCell.SetTextColor(tcell.ColorTurquoise)
		ol.SetCell(nextRow, 0, timestampCell)

		// Format the company
		company := strings.TrimSpace(offer.Company)
		companyCell := tview.NewTableCell(company)
		companyCell.SetTextColor(tcell.ColorGreen)
		ol.SetCell(nextRow, 1, companyCell)

		// Format the position
		position := strings.TrimSpace(offer.Position)
		positionCell := tview.NewTableCell(position)
		positionCell.SetExpansion(1)
		ol.SetCell(nextRow, 2, positionCell)

		// Put a backing ID so that we can refer to this offer later.
		ol.backingOfferIds[nextRow] = offer.ID
		nextRow++
	}
}
