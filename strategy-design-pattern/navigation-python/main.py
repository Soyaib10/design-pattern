from dataclasses import dataclass
from typing import List


@dataclass(frozen=True)
class Location:
    name: str
    lat: float
    lng: float

    def __str__(self) -> str:
        return f"{self.name} ({self.lat:.4f}, {self.lng:.4f})"


@dataclass
class RouteResult:
    mode: str
    checkpoints: List[Location]
    distance_km: float
    duration_min: float

    def __str__(self) -> str:
        return (
            f"[{self.mode}] {len(self.checkpoints)} checkpoint(s) | "
            f"distance: {self.distance_km:.1f} km | "
            f"time: {self.duration_min:.0f} min"
        )


if __name__ == "__main__":
    home = Location(name="Home", lat=23.7808, lng=90.3667)
    airport = Location(name="Airport", lat=23.8433, lng=90.4079)

    print("Start point:", home)
    print("End point:", airport)

    sample_result = RouteResult(
        mode="car",
        checkpoints=[home, airport],
        distance_km=12.0,
        duration_min=18.0,
    )
    print("\nResult:", sample_result)
