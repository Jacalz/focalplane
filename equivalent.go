package main

import (
	"cmp"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func equivalentView() fyne.CanvasObject {
	focal := &widget.Entry{PlaceHolder: "mm", Validator: uintValidator}
	aperture := &widget.Entry{PlaceHolder: "f-stops", Validator: floatValidator}
	sensor := &widget.Select{Options: sensors[:]}

	text := &widget.Label{Text: "Equivalent on full frame:", TextStyle: fyne.TextStyle{Bold: true}}
	value := &widget.Label{}
	data := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Focal length", Widget: focal, HintText: "Given in millimeters."},
			{Text: "Aperture", Widget: aperture, HintText: "Given in f-stops."},
			{Text: "Sensor type", Widget: sensor, HintText: "Digital and analog formats."},
		},
	}
	recalculateDOF := func(_ string) {
		focallengh, errF := strconv.ParseUint(focal.Text, 10, 64)
		fstop, errA := strconv.ParseFloat(aperture.Text, 64)
		selection := sensor.SelectedIndex()
		if cmp.Or(errF, errA) != nil || selection < 0 {
			value.SetText("")
			return
		}

		crop := cropFactor[selection]
		value.SetText(fmt.Sprintf("%.2f mm, f %.1f", float64(focallengh)*crop, fstop*crop))
	}

	// Hook up widgets to recalculate when something changes.
	focal.OnChanged = recalculateDOF
	aperture.OnChanged = recalculateDOF
	sensor.OnChanged = recalculateDOF
	return container.NewBorder(nil, container.NewHBox(text, value), nil, nil, data)

}
