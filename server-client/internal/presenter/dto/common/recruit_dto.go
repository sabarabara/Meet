package common

import (
	"time"

	"github.com/google/uuid"
)

type RecruitDTO struct {
	recruitID   uuid.UUID
	userID      uuid.UUID
	area        string
	imgURL      string
	man         uint16
	woman       uint16
	vacantMan   uint16
	vacantWoman uint16
	comment     string
	date        time.Time
	latitude    float64
	longitude   float64
}

func NewRecruitDTO(recruitID, userID uuid.UUID, area, imgURL string, man, woman, vacantMan, vacantWoman uint16, comment string, date time.Time, latitude, longitude float64) *RecruitDTO {
	return &RecruitDTO{
		recruitID:   recruitID,
		userID:      userID,
		area:        area,
		imgURL:      imgURL,
		man:         man,
		woman:       woman,
		vacantMan:   vacantMan,
		vacantWoman: vacantWoman,
		comment:     comment,
		date:        date,
		latitude:    latitude,
		longitude:   longitude,
	}
}
func (r *RecruitDTO) RecruitID() uuid.UUID {
	return r.recruitID
}
func (r *RecruitDTO) UserID() uuid.UUID {
	return r.userID
}
func (r *RecruitDTO) Area() string {
	return r.area
}
func (r *RecruitDTO) ImgURL() string {
	return r.imgURL
}
func (r *RecruitDTO) Man() uint16 {
	return r.man
}
func (r *RecruitDTO) Woman() uint16 {
	return r.woman
}
func (r *RecruitDTO) VacantMan() uint16 {
	return r.vacantMan
}
func (r *RecruitDTO) VacantWoman() uint16 {
	return r.vacantWoman
}
func (r *RecruitDTO) Comment() string {
	return r.comment
}
func (r *RecruitDTO) Date() time.Time {
	return r.date
}

func (r *RecruitDTO) Latitude() float64 {
	return r.latitude
}

func (r *RecruitDTO) Longitude() float64 {
	return r.longitude
}
