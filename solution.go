package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import "math"


type SidesCount int

const (
	SidesTriangle SidesCount = 3
	SidesCircle   SidesCount = 0
	SidesSquare   SidesCount = 4
)

func CalcSquare(sideLen float64, sidesNum SidesCount) float64 {

	if sidesNum == SidesCircle {
		return math.Pi * sideLen * sideLen
	}

	if sidesNum == SidesTriangle {
		return (math.Sqrt(3) / 4) * sideLen * sideLen
	}

	if sidesNum == SidesSquare {
		return sideLen * sideLen
	}

	return 0
}
