package main

import "fmt"

// This is still the "bad" version.
// It is intentionally written this way
// so that the problem becomes obvious.

func buildRoute(mode string, start string, end string) {
	if mode == "car" {
		// Calculation for driving
		distanceKm := 12.0
		speedKmph := 40.0
		durationMin := (distanceKm / speedKmph) * 60
		fmt.Printf("Car Route: %s to %s | Distance: %.1f km | Time: %.0f minutes\n",
			start, end, distanceKm, durationMin)
	} else if mode == "walk" {
		// Calculation for walking
		distanceKm := 3.0 // Walking route is shorter because of shortcuts
		speedKmph := 5.0
		durationMin := (distanceKm / speedKmph) * 60
		fmt.Printf("Walking Route: %s to %s | Distance: %.1f km | Time: %.0f minutes\n",
			start, end, distanceKm, durationMin)
	} else if mode == "bus" {
		// Calculation for bus travel, including stop delays
		distanceKm := 14.0
		speedKmph := 25.0
		numStops := 5
		stopDelayMin := 2.0
		durationMin := (distanceKm/speedKmph)*60 + float64(numStops)*stopDelayMin
		fmt.Printf("Bus Route: %s to %s | Distance: %.1f km | Time: %.0f minutes (including stops)\n",
			start, end, distanceKm, durationMin)

	} else {
		fmt.Println("This route type is not supported:", mode)
	}
}

func main() {
	buildRoute("car", "Home", "Airport")
	buildRoute("walk", "Home", "Airport")
	buildRoute("bus", "Home", "Airport")
}
