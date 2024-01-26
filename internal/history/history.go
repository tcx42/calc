package history

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var History = [][]string{{"a", "b"}, {"c", "d"}}

func HistoryTable() *widget.Table {
	return widget.NewTable(
		func() (rows int, cols int) {
			return len(History), len(History[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("______________")
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(History[tci.Row][tci.Col])
		},
	)
}
