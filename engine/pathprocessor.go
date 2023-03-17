package engine

// 1. In a simple case the start of the itinerary does not have the matching 'end' destination;
// and the itinerary final destination does not have the matching 'start' within the input path segments.
// Ex.[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]], "SFO" does not have a matching 'end' in any [start, end] segments;
// "EWR" does not have any matching 'start', while other intermediate destinations such as "ATL", can be found as both 'start' and 'end' of
// among the segments of the itinerary.
// 2. In case of a true cycle all start locations will have matching destination locations. This is how we can detect a cycle in the input
// Ex. [["ATL", "EWR"], ["SFO", "ATL"], ["EWR", "SFO"]]
// To disambiguate a cycle and find the starting point, filght times are needed.
// Without the time information, we can start the cycle from any point on the itinerary.
// This version of API will return a 'BadDataErr' with a detailed logged message if a cycle is detected.
// 3. If there were cycles during the trip, but it eventually arrived to another destination,
// the beginning and the end of the trip are returned like in the simple case
// 4. If multiple 'start' and/or 'end' locations detected in the input data, 'BadDataErr' with detailed message is returned

// 'process' returns start & end of the trip data, or it returns nil and error if problem with input detected
func process(paths [][]string) ([]string, error) {
	combined := make(map[string]int, 0)
	for _, path := range paths {
		combined[path[0]]++
		combined[path[1]]--
	}
	// The start will have count == 1
	// The destination will have count == -1
	begin := ""
	end := ""
	for key, count := range combined {
		if count == 0 {
			continue
		}
		if count == 1 && begin == "" {
			begin = key
		} else if count == -1 && end == "" {
			end = key
		} else if count > 1 && begin == "" {
			return nil, newBadDataErr(multipleStartsEndsErr, []string{key}, nil)
		} else if count < -1 && end == "" {
			return nil, newBadDataErr(multipleStartsEndsErr, nil, []string{key})
		} else if count >= 1 && begin != "" {
			return nil, newBadDataErr(multipleStartsEndsErr, []string{begin, key}, nil)
		} else if count <= -1 && end != "" {
			return nil, newBadDataErr(multipleStartsEndsErr, nil, []string{end, key})
		}
	}
	if begin == "" && end == "" {
		return nil, newBadDataErr(cyclicalErr, nil, nil)
	}
	return []string{begin, end}, nil
}
