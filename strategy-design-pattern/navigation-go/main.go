package main

import "fmt"

type Location struct {
	Name string
	Lat  float32
	Lng  float32
}

func (l Location) String() string {
	return fmt.Sprintf("%s (%.4f, %.4f)", l.Name, l.Lat, l.Lng)
}

type RouteResult struct {
	Mode        string
	Checkpoints []Location
	DistanceKm  float32
	DurationMin float32
}

func (r RouteResult) String() string {
	return fmt.Sprintf(
		"[%s] %d checkpoint(s) | distance: %.1f km | time: %.0f min",
		r.Mode, len(r.Checkpoints), r.DistanceKm, r.DurationMin,
	)
}

func main() {
	home := Location{
		Name: "Home",
		Lat:  32.324,
		Lng:  90.234,
	}

	airport := Location{
		Name: "airport",
		Lat:  31.324,
		Lng:  40.234,
	}

	fmt.Println("Start point:", home)
	fmt.Println("End point:", airport)

	sampleResult := RouteResult{
		Mode: "car",
		Checkpoints: []Location{
			home,
			airport,
		},
		DistanceKm: 32.3,
		DurationMin: 234.3,
	}

	fmt.Println("\nResult: ", sampleResult)
}
