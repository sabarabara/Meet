package recruitlocation

import (
	"errors"

	"github.com/google/uuid"
)

const (
	errorEmptyRecruitLocationID = "募集位置IDが空になっています"
	errorEmptyRecruitID         = "募集IDが空になっています"
	errorEmptyLatitude          = "緯度が空になっています"
	errorEmptyLongitude         = "経度が空になっています"
)

type RecruitLocation struct {
	recruitlocationid uuid.UUID
	recruitid         uuid.UUID
	latitude          float64
	longitude         float64
}

func NewRecruitLocation(recruitlocationid uuid.UUID, recruitid uuid.UUID, latitude float64, longitude float64) (RecruitLocation, error) {
	if recruitlocationid == uuid.Nil {
		return RecruitLocation{}, errors.New(errorEmptyRecruitLocationID)
	}
	if recruitid == uuid.Nil {
		return RecruitLocation{}, errors.New(errorEmptyRecruitID)
	}
	if latitude == 0 {
		return RecruitLocation{}, errors.New(errorEmptyLatitude)
	}
	if longitude == 0 {
		return RecruitLocation{}, errors.New(errorEmptyLongitude)
	}
	return RecruitLocation{
		recruitlocationid: recruitlocationid,
		recruitid:         recruitid,
		latitude:          latitude,
		longitude:         longitude,
	}, nil
}

func (rl RecruitLocation) Recruitlocationid() uuid.UUID {
	return rl.recruitlocationid
}
func (rl RecruitLocation) Recruitid() uuid.UUID {
	return rl.recruitid
}
func (rl RecruitLocation) Latitude() float64 {
	return rl.latitude
}
func (rl RecruitLocation) Longitude() float64 {
	return rl.longitude
}
