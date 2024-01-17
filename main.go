package main

import (
	"fmt"
	"scicalc/pkg/eval"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("SciCalc")

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
		buttonsGrid(input, result),
	)
	return box2
}

func buttonsGrid(input *widget.Entry, output *widget.Label) *fyne.Container {
	var mainButtons = []fyne.CanvasObject{
		widget.NewButton(" 7 ", func() { input.Append("7") }),
		widget.NewButton(" 8 ", func() { input.Append("8") }),
		widget.NewButton(" 9 ", func() { input.Append("9") }),
		widget.NewButton(" C ", func() {
			if input.Text == "" {
				output.SetText("")
			}
			input.SetText("")
		}),
		widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
			if len(input.Text) > 0 {
				input.SetText(input.Text[:len(input.Text)-1])
			}
		}),
		widget.NewButton(" 4 ", func() { input.Append("4") }),
		widget.NewButton(" 5 ", func() { input.Append("5") }),
		widget.NewButton(" 6 ", func() { input.Append("6") }),
		widget.NewButtonWithIcon("", theme.CancelIcon(), func() { input.Append("*") }),
		widget.NewButton(" / ", func() { input.Append("/") }),
		widget.NewButton(" 1 ", func() { input.Append("1") }),
		widget.NewButton(" 2 ", func() { input.Append("2") }),
		widget.NewButton(" 3 ", func() { input.Append("3") }),
		widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() { input.Append("+") }),
		widget.NewButtonWithIcon("", theme.ContentRemoveIcon(), func() { input.Append("-") }),
		widget.NewButton(" E ", func() { input.Append("E") }),
		widget.NewButton(" 0 ", func() { input.Append("0") }),
		widget.NewButton(" . ", func() { input.Append(".") }),
		widget.NewButton("Ans", func() { input.Append(output.Text) }),
		widget.NewButton(" = ", func() {
			res, err := eval.Eval(input.Text)
			if err == nil {
				output.SetText(fmt.Sprint(res))
			} else {
				output.SetText(fmt.Sprint(err))
			}
		}),
	}
	var secButtons = []fyne.CanvasObject{
		widget.NewButton("", func() {}),
		widget.NewButton("", func() {}),
		widget.NewButton("", func() {}),
		widget.NewButton("x²", func() { input.Append("^2") }),
		widget.NewButton("(", func() { input.Append("(") }),
		widget.NewButton(")", func() { input.Append(")") }),
	}
	mainGrid := container.NewGridWithColumns(5, mainButtons...)
	secGrid := container.NewGridWithColumns(6, secButtons...)
	box := container.NewVBox(secGrid, mainGrid)
	return box
}
