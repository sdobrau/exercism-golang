package spaceage

type Planet string

const MERCURY_PERIOD = 0.2408467
const VENUS_PERIOD = 0.61519726
const EARTH_PERIOD = 1.0
const MARS_PERIOD = 1.8808158
const JUPITER_PERIOD = 11.862615
const SATURN_PERIOD = 29.447498
const URANUS_PERIOD = 84.016846
const NEPTUNE_PERIOD = 164.79132

func Age(seconds float64, planet Planet) float64 {
	planet_age := 0.0
	switch planet {
	case "Mercury": planet_age = seconds / (31557600.0 * MERCURY_PERIOD)
	case "Venus": planet_age = seconds / (31557600.0 * VENUS_PERIOD)
	case "Mars": planet_age = seconds / (31557600.0 * MARS_PERIOD)
	case "Earth": planet_age = seconds / (31557600.0 * EARTH_PERIOD)
	case "Jupiter": planet_age = seconds / (31557600.0 * JUPITER_PERIOD)
	case "Saturn": planet_age = seconds / (31557600.0 * SATURN_PERIOD)
	case "Uranus": planet_age = seconds / (31557600.0 * URANUS_PERIOD)
	case "Neptune": planet_age = seconds / (31557600.0 * NEPTUNE_PERIOD)
	default: return -1
	}
	return planet_age

}
