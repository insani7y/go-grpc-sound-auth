package fileHandler

import (
	"math"
	"math/cmplx"
)

func WindowTransform(num float64, ind int) float64 {
	return num * (0.53836 - 0.46164 * math.Cos(float64(ind) * 2 * math.Pi / (FramesCount - 1)))
}

func HerzToMel(value float64) float64 {
	return 1127 * math.Log(1 + value/700)
}

func CalculateAmplitude(frameNumber int, frame []float64) []float64{
	length := float64(len(frame))
	for n, value := range frame {
		complexData :=  complex(0, -2 * float64(frameNumber) * math.Pi * float64(n) / length)
		complexExp := cmplx.Exp(complexData)
		frame[n] = cmplx.Abs(complex(value, 0) * complexExp)
	}

	return frame
}