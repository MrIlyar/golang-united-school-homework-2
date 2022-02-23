package square

import (
	"math"
)

type numberOfSides int

const SidesTriangle numberOfSides = 3
const SidesSquare numberOfSides = 4
const SidesCircle numberOfSides = 0

func CalcSquare(sideLen float64, sidesNum numberOfSides) float64 {
	switch sidesNum {
	case SidesTriangle:
		return sideLen * sideLen * math.Sqrt(3) / 4
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return sideLen * sideLen * math.Pi
	default:
		return 0
	}
}
