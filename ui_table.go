package main

import (
	"github.com/rivo/tview"
)

type HeaderTable struct {
	*tview.Table
}

func NewHeaderTable() *HeaderTable {
	return &HeaderTable{Table: tview.NewTable()}
}

func (ht *HeaderTable) AddRow(label string, value *tview.TableCell) {
	rowCount := ht.GetRowCount()
	ht.SetCellSimple(rowCount, 0, label)
	ht.SetCell(rowCount, 1, value)
}
