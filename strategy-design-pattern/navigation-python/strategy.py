from abc import ABC, abstractmethod

from models import Location, RouteResult


class RouteStrategy(ABC):
    @abstractmethod
    def build_route(self, start: Location, end: Location) -> RouteResult:
        pass


class CarRouteStrategy(RouteStrategy):
    def build_route(self, start: Location, end: Location) -> RouteResult:
        distance = 12.2
        speed_kmh = 39.3
        duration = (distance / speed_kmh) * 60

        return RouteResult(
            mode="Car",
            checkpoints=[start, end],
            distance_km=distance,
            duration_min=duration,
        )


class WalkRouteStrategy(RouteStrategy):
    def build_route(self, start: Location, end: Location) -> RouteResult:
        distance = 3.5
        speed_kmh = 5.0
        duration = (distance / speed_kmh) * 60

        return RouteResult(
            mode="Walk",
            checkpoints=[start, end],
            distance_km=distance,
            duration_min=duration,
        )

class BusRouteStrategy(RouteStrategy):
    def __init__(self, num_stops: int = 5, stop_delay_min: float = 2.0):
        self.num_stops = num_stops
        self.stop_delay_min = stop_delay_min

    def build_route(self, start: Location, end: Location) -> RouteResult:
        distance = 14.0
        speed_kmh = 25.0
        travel_time = (distance / speed_kmh) * 60
        total_duration = travel_time + (self.num_stops * self.stop_delay_min)

        checkpoints = [start]
        for i in range(1, self.num_stops + 1):
            lat_offset = (end.lat - start.lat) * (i / (self.num_stops + 1))
            lng_offset = (end.lng - start.lng) * (i / (self.num_stops + 1))
            checkpoints.append(
                Location(
                    name=f"Bus Stop {i}",
                    lat=start.lat + lat_offset,
                    lng=start.lng + lng_offset,
                )
            )
        checkpoints.append(end)

        return RouteResult(
            mode=f"Bus ({self.num_stops} stops)",
            checkpoints=checkpoints,
            distance_km=distance,
            duration_min=total_duration,
        )
