package fileHandler

import (
	"math"
)

type FileHandler struct{}

func New() *FileHandler {
	return &FileHandler{}
}

func (f *FileHandler) DetermineAmplitudeValues(fileBytes []byte) ([]float64, error) {
	var MSB, LSB float64

	data := make([]float64, len(fileBytes) / 2)

	for i := 0; i < len(fileBytes) - 1; i += 2 {
		MSB = float64(fileBytes[i])
		LSB = float64(fileBytes[i + 1])
		data[i/2] = ((MSB * math.Pow(2, 8)) + LSB) / math.Pow(2, 16)
	}

	return data, nil
}

func (f *FileHandler) FrameCut(data []float64) [][]float64 {

	res := make([][]float64, FramesCount)

	shortLength := len(data) / FramesCount

	for c := 0; c < FramesCount; c++ {
		for i := 0; i < shortLength; i ++ {
			res[c] = append(res[c], WindowTransform(data[(c * shortLength) + i], i))
		}
	}

	return res
}

func (f *FileHandler) FourierTransform(frames [][]float64) [][]float64 {
	for n, frame := range frames {
		frames[n] = CalculateAmplitude(n, frame)
	}

	return frames
}

func (f *FileHandler) ToMelScale(amplitudes [][]float64) [][]float64 {

	for i, amplitude := range amplitudes {
		for j, value := range amplitude {
			amplitudes[i][j] = HerzToMel(value)
		}
	}

	return amplitudes
}

func (f *FileHandler) GetMelFeaturesArr(amplitudes [][]float64) []float64 {
	var features []float64
	var summs []float64

	for _, amplitudeArr := range amplitudes {
		var res float64 = 0
		for _, melValue := range amplitudeArr {
			if melValue > 0 {
				res += math.Log(melValue)
			}
		}
		summs = append(summs, res)
	}

	for n := 1; n <= MelCoeffCount; n++ {
		finalValue := summs[n-1] * math.Pi / MelCoeffCount
		features = append(features, finalValue)
	}

	return features
}
