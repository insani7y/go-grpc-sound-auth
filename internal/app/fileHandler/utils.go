package fileHandler

import "math"

func WindowTransform(num float64, ind int) float64 {
	return num * (0.53836 - 0.46164 * math.Cos(float64(ind) * 2 * math.Pi / (FramesCount - 1)))
}
