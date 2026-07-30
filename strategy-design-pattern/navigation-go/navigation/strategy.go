package main

type RouteStrategy interface {
	BuildRoute(start, end Location) RouteResult
}

type CarRouteStrategy struct{}

func (c CarRouteStrategy) BuildRoute(start, end Location) RouteResult {
	distanceKm := 12.0
	speedKmph := 40.0
	durationMin := (distanceKm / speedKmph) * 60

	return RouteResult{
		Mode:        "car",
		Checkpoints: []Location{start, end},
		DistanceKm:  distanceKm,
		DurationMin: durationMin,
	}
}


type WalkRouteStrategy struct{}

func (w WalkRouteStrategy) BuildRoute(start, end Location) RouteResult {
	distanceKm := 3.0
	speedKmph := 5.0
	durationMin := (distanceKm / speedKmph) * 60

	return RouteResult{
		Mode:        "walk",
		Checkpoints: []Location{start, end},
		DistanceKm:  distanceKm,
		DurationMin: durationMin,
	}
}


type PublicTransportRouteStrategy struct {
	NumStops     int
	StopDelayMin float64
}

func (p PublicTransportRouteStrategy) BuildRoute(start, end Location) RouteResult {
	distanceKm := 14.0
	speedKmph := 25.0
	travelMin := (distanceKm / speedKmph) * 60
	totalStopDelay := float64(p.NumStops) * p.StopDelayMin
	durationMin := travelMin + totalStopDelay

	return RouteResult{
		Mode:        "bus",
		Checkpoints: []Location{start, end},
		DistanceKm:  distanceKm,
		DurationMin: durationMin,
	}
}