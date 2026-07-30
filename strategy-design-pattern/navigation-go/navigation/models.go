package main

import "fmt"

type Location struct {
	Name string
	Lat  float64
	Lng  float64
}

func (l Location) String() string {
	return fmt.Sprintf("%s (%.4f, %.4f)", l.Name, l.Lat, l.Lng)
}

type RouteResult struct {
	Mode        string
	Checkpoints []Location
	DistanceKm  float64
	DurationMin float64
}

func (r RouteResult) String() string {
	return fmt.Sprintf(
		"[%s] %d checkpoint(s) | distance: %.1f km | time: %.0f min",
		r.Mode, len(r.Checkpoints), r.DistanceKm, r.DurationMin,
	)
}