package main

var sensors = [...]string{
	// Small format
	"Fullframe (35mm)",
	"APS-H (Canon)",
	"APS-C (Generic)",
	"APS-C (Canon)",
	"Micro Four Thirds",
	"1\"",
	// Medium format
	"6x9",
	"6x7",
	"6x6",
	"645 (6x4.5)",
	// Large format
	"8x10",
	"5x7",
	"4x5",
}

// Sensor name to circle of confusion in meters.
// Based on table from https://en.wikipedia.org/wiki/Circle_of_confusion.
var circleOfConfusion = [len(sensors)]float64{
	// Small format
	0.029e-3,
	0.023e-3,
	0.019e-3,
	0.018e-3,
	0.015e-3,
	0.011e-3,

	// Medium format
	0.067e-3,
	0.059e-3,
	0.053e-3,
	0.047e-3,

	// Large format
	0.22e-3,
	0.15e-3,
	0.11e-3,
}

var cropFactor = [len(sensors)]float64{
	// Small format
	1.0,
	1.3,
	1.5,
	1.6,
	2.0,
	2.7,

	// Medium format
	0.43,
	0.5,
	0.55,
	0.62,

	// Large format
	0.13,
	0.20,
	0.27,
}
