package main

import (
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
	dofText := &widget.Label{Text: "Depth of field:", TextStyle: fyne.TextStyle{Bold: true}}
	dofValue := &widget.Label{}
	dofData := &widget.Form{
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
		focallengh, _ := strconv.ParseUint(focal.Text, 10, 64)
		distance, _ := strconv.ParseFloat(distance.Text, 64)
		fstop, _ := strconv.ParseFloat(aperture.Text, 64)
		circle := sensorToCoC[sensor.Selected]

		dof := depthOfField(float64(focallengh)/1000, distance, fstop, circle)
		dofValue.SetText(strconv.FormatFloat(dof, 'f', 6, 64) + " m")
	}

	// Hook up widgets to recalculate when something changes.
	focal.OnChanged = recalculateDOF
	distance.OnChanged = recalculateDOF
	aperture.OnChanged = recalculateDOF
	sensor.OnChanged = recalculateDOF
	return container.NewBorder(nil, container.NewHBox(dofText, dofValue), nil, nil, dofData)
}

// Based on algorithm from https://en.wikipedia.org/wiki/Depth_of_field.
func depthOfField(focallength, distance, aperture, circle float64) float64 {
	return 2 * distance * distance * aperture * circle / (focallength * focallength)
}
