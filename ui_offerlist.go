package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"strings"
)

type OfferList struct {
	*tview.Table
	backingOffers   []Offer
	backingOfferIds map[int]int
	filterFunc      func(offer Offer) bool
}

func NewOfferList() *OfferList {
	offerList := new(OfferList)
	offerList.Table = tview.NewTable()
	offerList.SetSelectable(true, false)
	offerList.filterFunc = nil
	return offerList
}

func (ol *OfferList) SetFilterFunc(filter func(offer Offer) bool) {
	ol.filterFunc = filter
	ol.reloadTable()
}

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
