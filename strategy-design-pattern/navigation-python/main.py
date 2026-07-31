from models import Location
from navigator import Navigator
from strategy import BusRouteStrategy, CarRouteStrategy, WalkRouteStrategy

if __name__ == "__main__":
    home = Location(name="Home", lat=23.7808, lng=90.3667)
    airport = Location(name="Airport", lat=23.8433, lng=90.4079)

    nav = Navigator(CarRouteStrategy())
    print("=== Car Route ===")
    print(nav.build_route(home, airport))

    nav.set_strategy(WalkRouteStrategy())
    print("\n=== Walk Route (Runtime Switch) ===")
    print(nav.build_route(home, airport))

    nav.set_strategy(BusRouteStrategy(num_stops=6, stop_delay_min=1.5))
    print("\n=== Bus Route ===")
    print(nav.build_route(home, airport))
