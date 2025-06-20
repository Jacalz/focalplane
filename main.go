package main

import (
	_ "embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/validation"
)

var uintValidator = validation.NewRegexp(`^\d+$`, "Must be a positive integer type value.")
var floatValidator = validation.NewRegexp(`\d+(\.\d+)?$`, "Must be a valid decimal number.")

//go:embed img/icon.png
var icon []byte

func main() {
	a := app.NewWithID("io.github.jacalz.focalplane")
	a.SetIcon(fyne.NewStaticResource("icon.png", icon))
	w := a.NewWindow("Focalplane")

	w.SetContent(dofView())
	w.ShowAndRun()
}
