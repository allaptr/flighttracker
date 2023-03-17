package engine

type FlightTracker interface {
	Track(itinerary [][]string) ([]string, error)
}

func Track(itinerary [][]string) ([]string, error) {
	return process(itinerary)
}
