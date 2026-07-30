package main

import "fmt"

type Navigator struct {
	strategy RouteStrategy
}

func NewNavigator(strategy RouteStrategy) *Navigator {
	return &Navigator{strategy: strategy}
}

func (n *Navigator) SetStrategy(strategy RouteStrategy) {
	n.strategy = strategy
}

func (n *Navigator) BuildRoute(start, end Location) RouteResult {
	return n.strategy.BuildRoute(start, end)
}

func main() {
	home := Location{Name: "Home", Lat: 23.7808, Lng: 90.3667}
	airport := Location{Name: "Airport", Lat: 23.8433, Lng: 90.4079}

	nav := NewNavigator(CarRouteStrategy{})
	fmt.Println("User picked: car")
	fmt.Println(nav.BuildRoute(home, airport))

	nav.SetStrategy(WalkRouteStrategy{})
	fmt.Println("\nUser picked: walk")
	fmt.Println(nav.BuildRoute(home, airport))

	nav.SetStrategy(PublicTransportRouteStrategy{NumStops: 5, StopDelayMin: 2.0})
	fmt.Println("\nUser picked: bus")
	fmt.Println(nav.BuildRoute(home, airport))
}
