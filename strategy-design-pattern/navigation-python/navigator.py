from models import Location, RouteResult
from strategy import RouteStrategy


class Navigator:
    def __init__(self, strategy: RouteStrategy):
        self._strategy = strategy

    def set_strategy(self, strategy: RouteStrategy) -> None:
        self._strategy = strategy

    def build_route(self, start: Location, end: Location) -> RouteResult:
        return self._strategy.build_route(start, end)
