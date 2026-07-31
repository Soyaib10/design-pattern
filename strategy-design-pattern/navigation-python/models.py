def build_route(mode: str, start: str, end: str) -> None:
    # This is still the "bad" version — written intentionally this way
    # so we can see the problem with our own eyes.
    if mode == "car":
        distance_km = 12.0
        speed_kmph = 40.0
        duration_min = (distance_km / speed_kmph) * 60
        print(f"Car route: {start} -> {end} | distance: {distance_km:.1f} km | time: {duration_min:.0f} min")

    elif mode == "walk":
        distance_km = 3.0  # walking path uses a shortcut, so shorter
        speed_kmph = 5.0
        duration_min = (distance_km / speed_kmph) * 60
        print(f"Walk route: {start} -> {end} | distance: {distance_km:.1f} km | time: {duration_min:.0f} min")

    elif mode == "bus":
        distance_km = 14.0
        speed_kmph = 25.0
        num_stops = 5
        stop_delay_min = 2.0
        duration_min = (distance_km / speed_kmph) * 60 + num_stops * stop_delay_min
        print(f"Bus route: {start} -> {end} | distance: {distance_km:.1f} km | time: {duration_min:.0f} min (incl. stops)")

    else:
        print("Unsupported route mode:", mode)


if __name__ == "__main__":
    build_route("car", "Home", "Airport")
    build_route("walk", "Home", "Airport")
    build_route("bus", "Home", "Airport")