package main

import (
	"scicalc/internal/buttons"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("Calc")

	window.SetContent(makeUI())
	window.ShowAndRun()
}

func makeUI() fyne.CanvasObject {
	input := widget.NewEntry()
	result := widget.NewLabel("")
	box1 := container.NewVBox(
		input,
		result,
	)
	box2 := container.NewVSplit(
		box1,
		buttons.ButtonsGrid(input, result),
	)
	return box2
}
