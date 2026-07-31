from typing import List


class Location:
    def __init__(self, name: str, lat: float, lng: float):
        self.name = name
        self.lat = lat
        self.lng = lng

    def __str__(self) -> str:
        return f"{self.name} ({self.lat:.4f}, {self.lng:.4f})"


class RouteResult:
    def __init__(self, mode: str, checkpoints: List["Location"], distance_km: float, duration_min: float,
    ):
        self.mode = mode
        self.checkpoints = checkpoints
        self.distance_km = distance_km
        self.duration_min = duration_min

    def __str__(self) -> str:
        return (
            f"[{self.mode}] {len(self.checkpoints)} checkpoint(s) | "
            f"distance: {self.distance_km:.1f} km | "
            f"time: {self.duration_min:.0f} min"
        )
