package auth

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"math"
)

type UserAuthData struct {
	UserId int
	Features DataFeatures
	AuthDataId []uint8
}

type DataFeatures struct {
	MFCC [][]float64
}

type UserSoundDifferenceMap struct {
	data map[int]float64
}

func NewMap() *UserSoundDifferenceMap {
	return &UserSoundDifferenceMap{
		data: make(map[int]float64),
	}
}

func (u *UserSoundDifferenceMap) UserMin() int {
	min := math.MaxFloat64
	var minUserId int
	for userId, difference := range u.data {
		if min > difference {
			min = difference
			minUserId = userId
		}
	}
	return minUserId
}

func (u *UserSoundDifferenceMap) AddToKey(userId int, value float64) {
	if _, ok := u.data[userId]; !ok {
		u.data[userId] = value
	} else {
		u.data[userId] += value
	}
}

func (d DataFeatures) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (d *DataFeatures) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &d)
}

func (d *DataFeatures) CalculateDifference(mfcc [][]float64) float64 {
	var res float64

	for i := 0; i < len(mfcc); i++ {
		var diff []float64
		for j := 0; j < len(mfcc[i]); j++ {
			diff = append(diff, d.MFCC[i][j] - mfcc[i][j])
		}
		res += CalculateVecAbs(diff)
	}

	return res
}