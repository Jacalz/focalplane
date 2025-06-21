package main

import (
	"cmp"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func dofView() fyne.CanvasObject {
	// Camera and lens parameters.
	focal := &widget.Entry{PlaceHolder: "mm", Validator: uintValidator}
	distance := &widget.Entry{PlaceHolder: "m", Validator: floatValidator}
	aperture := &widget.Entry{PlaceHolder: "f-stops", Validator: floatValidator}
	sensor := &widget.Select{Options: sensors[:]}

	// Building blocks for the user interface.
	text := &widget.Label{Text: "Depth of field:", TextStyle: fyne.TextStyle{Bold: true}}
	value := &widget.Label{}
	data := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Focal length", Widget: focal, HintText: "Given in millimeters."},
			{Text: "Distance to subject", Widget: distance, HintText: "Given in meters."},
			{Text: "Aperture", Widget: aperture, HintText: "Given in f-stops."},
			{Text: "Sensor type", Widget: sensor, HintText: "Digital and analog formats."},
		},
	}

	// Calculate the depth of field from values above.
	// The validators guarantee that we don't get invalid inputs.
	recalculateDOF := func(_ string) {
		focallengh, errF := strconv.ParseUint(focal.Text, 10, 64)
		distance, errD := strconv.ParseFloat(distance.Text, 64)
		fstop, errA := strconv.ParseFloat(aperture.Text, 64)
		selection := sensor.SelectedIndex()
		if cmp.Or(errF, errD, errA) != nil || selection < 0 {
			value.SetText("")
			return
		}

		dof := depthOfField(float64(focallengh)/1000, distance, fstop, circleOfConfusion[selection])
		value.SetText(strconv.FormatFloat(dof, 'f', 6, 64) + " m")
	}

	// Hook up widgets to recalculate when something changes.
	focal.OnChanged = recalculateDOF
	distance.OnChanged = recalculateDOF
	aperture.OnChanged = recalculateDOF
	sensor.OnChanged = recalculateDOF
	return container.NewBorder(nil, container.NewHBox(text, value), nil, nil, data)
}

// Based on algorithm from https://en.wikipedia.org/wiki/Depth_of_field.
func depthOfField(focallength, distance, aperture, circle float64) float64 {
	return 2 * distance * distance * aperture * circle / (focallength * focallength)
}
