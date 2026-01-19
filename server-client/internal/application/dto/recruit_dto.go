package dto

import (
	"time"

	"github.com/google/uuid"
)

type RecruitDTO struct {
	recruitid    *uuid.UUID
	userid       uuid.UUID
	area         string
	imgurl       string
	man          uint16
	woman        uint16
	vacant_man   uint16
	vacant_woman uint16
	commnt       string
	date         time.Time
}

func NewRecruitDTO(
	recruitid *uuid.UUID,
	userid uuid.UUID,
	area string,
	imgurl string,
	man uint16,
	woman uint16,
	vacant_man uint16,
	vacant_woman uint16,
	commnt string,
	date time.Time,
) RecruitDTO {
	return RecruitDTO{
		recruitid:    recruitid,
		userid:       userid,
		area:         area,
		imgurl:       imgurl,
		man:          man,
		woman:        woman,
		vacant_man:   vacant_man,
		vacant_woman: vacant_woman,
		commnt:       commnt,
		date:         date,
	}
}

func (r RecruitDTO) Recruitid() *uuid.UUID {
	return r.recruitid
}
func (r RecruitDTO) Userid() uuid.UUID {
	return r.userid
}
func (r RecruitDTO) Area() string {
	return r.area
}
func (r RecruitDTO) Imgurl() string {
	return r.imgurl
}
func (r RecruitDTO) Man() uint16 {
	return r.man
}
func (r RecruitDTO) Woman() uint16 {
	return r.woman
}
func (r RecruitDTO) VacantMan() uint16 {
	return r.vacant_man
}
func (r RecruitDTO) VacantWoman() uint16 {
	return r.vacant_woman
}
func (r RecruitDTO) Commnt() string {
	return r.commnt
}
func (r RecruitDTO) Date() time.Time {
	return r.date
}
