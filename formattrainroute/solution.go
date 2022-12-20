package formattrainroute

import "fmt"

func Format(stations []string) string {
	route := fmt.Sprintf("Train is calling at %s", stations[0])

	if len(stations) == 1 {
		return route
	}
	if len(stations) == 2 {
		return fmt.Sprintf("%s and %s", route, stations[1])
	}

	for i := 1; i <= len(stations)-2; i++ {
		route += fmt.Sprintf(", %s", stations[i])
	}
	route += fmt.Sprintf(" and %s", stations[len(stations)-1])
	return route
}
