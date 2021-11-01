package fileHandler

import "math"

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
