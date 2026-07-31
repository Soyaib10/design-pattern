package main

import "fmt"

func strategyFromMode(mode string) (RouteStrategy, error) {
	switch mode {
	case "car":
		return CarRouteStrategy{}, nil
	case "walk":
		return WalkRouteStrategy{}, nil
	case "bus":
		return PublicTransportRouteStrategy{NumStops: 5, StopDelayMin: 2.0}, nil
	case  "bicycle":
		return BicycleRouteStrategy{}, nil
	default:
		return nil, fmt.Errorf("unsupported mode: %q", mode)
	}
}

func main() {
	home := Location{Name: "Home", Lat: 23.7808, Lng: 90.3667}
	airport := Location{Name: "Airport", Lat: 23.8433, Lng: 90.4079}

	simulatedUserInputs := []string{"car", "walk", "bus", "scooter"}

	nav := NewNavigator(nil)

	for _, mode := range simulatedUserInputs {
		fmt.Printf("User selected: %q\n", mode)

		strategy, err := strategyFromMode(mode)
		if err != nil {
			fmt.Println("  ->", err)
			continue
		}

		nav.SetStrategy(strategy)
		result := nav.BuildRoute(home, airport)
		fmt.Println("  ->", result)
	}
}
