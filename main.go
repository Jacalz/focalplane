package main

import "fyne.io/fyne/v2/app"

func main() {
	a := app.NewWithID("io.github.jacalz.focalplane")
	w := a.NewWindow("Focalplane")

	w.ShowAndRun()
}
