// Package space which has a function to convert age accordingly
package space

type Planet string

// Age returns the age respective to the planet
func Age(seconds float64, planet Planet) float64 {
	earth_sec := 31557600
	switch planet {
	case "Mercury":
		return (seconds / (0.2408467 * float64(earth_sec)))
	case "Venus":
		return (seconds / (0.61519726 * float64(earth_sec)))
	case "Earth":
		return (seconds / float64(earth_sec))
	case "Mars":
		return (seconds / (1.8808158 * float64(earth_sec)))
	case "Jupiter":
		return (seconds / (11.862615 * float64(earth_sec)))
	case "Saturn":
		return (seconds / (29.447498 * float64(earth_sec)))
	case "Uranus":
		return (seconds / (84.016846 * float64(earth_sec)))
	case "Neptune":
		return (seconds / (164.79132 * float64(earth_sec)))
	default:
		return 0
	}
}
