package main

import (
	"github.com/rivo/tview"
)

// LocationsTable is a table widget used to present locations with offers.
type LocationsTable struct {
	// The backing table that presents the list of locations.
	*tview.Table

	// This map indicates which location is present at each row of the table.
	rowLocationIndex map[int]Location
}

// NewLocationsTable builds a table widget that can be used to display a list
// of locations where offers are available through the feed. Then, it's
// possible to view offers for that particular location.
func NewLocationsTable() *LocationsTable {
	table := &LocationsTable{
		Table:            tview.NewTable(),
		rowLocationIndex: make(map[int]Location),
	}

	// Populate the table with locations.
	nextRow := 0
	for key, location := range Locations {
		cell := tview.NewTableCell(location.title)
		cell.SetExpansion(1)
		table.SetCell(nextRow, 0, cell)
		table.rowLocationIndex[nextRow] = key
		nextRow++
	}

	table.SetSelectable(true, false)

	return table
}

// GetSelectedLocation converts the selected index of the table into the
// valid location constant that can be used by the context to fetch new
// job offers.
func (lt *LocationsTable) GetSelectedLocation() Location {
	selectedRow, _ := lt.GetSelection()
	return lt.rowLocationIndex[selectedRow]
}
