package dto

import "github.com/google/uuid"

type RecruitLocationDTO struct {
	recruitlocationid *uuid.UUID
	recruitid         uuid.UUID
	latitude          float64
	longitude         float64
}

func NewRecruitLocationDTO(recruitlocationid *uuid.UUID, recruitid uuid.UUID, latitude float64, longitude float64) RecruitLocationDTO {
	return RecruitLocationDTO{
		recruitlocationid: recruitlocationid,
		recruitid:         recruitid,
		latitude:          latitude,
		longitude:         longitude,
	}
}

func (r RecruitLocationDTO) Recruitlocationid() *uuid.UUID {
	return r.recruitlocationid
}
func (r RecruitLocationDTO) Recruitid() uuid.UUID {
	return r.recruitid
}
func (r RecruitLocationDTO) Latitude() float64 {
	return r.latitude
}
func (r RecruitLocationDTO) Longitude() float64 {
	return r.longitude
}
