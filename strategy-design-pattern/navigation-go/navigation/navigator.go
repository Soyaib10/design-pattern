package main

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
