package recruit

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	errorEmptyRecruitID = "募集IDが空になっています"
	errorEmptyUserID    = "ユーザーIDが空になっています"
	errorEmptyArea      = "募集エリアが空になっています"
	errorEmptyImgurl    = "画像URLが空になっています"
	errorEmptyDate      = "日付が空になっています"
)

type Recruit struct {
	recruitid    uuid.UUID
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

func NewRecruit(recruitid uuid.UUID, userid uuid.UUID, area string, imgurl string, man uint16, woman uint16, vacant_man uint16, vacant_woman uint16, commnt string, date time.Time) (Recruit, error) {
	if recruitid == uuid.Nil {
		return Recruit{}, errors.New(errorEmptyRecruitID)
	}
	if userid == uuid.Nil {
		return Recruit{}, errors.New(errorEmptyUserID)
	}
	if area == "" {
		return Recruit{}, errors.New(errorEmptyArea)
	}
	if imgurl == "" {
		return Recruit{}, errors.New(errorEmptyImgurl)
	}
	if date.IsZero() {
		return Recruit{}, errors.New(errorEmptyDate)
	}

	return Recruit{
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
	}, nil
}
func (r Recruit) Recruitid() uuid.UUID {
	return r.recruitid
}
func (r Recruit) Userid() uuid.UUID {
	return r.userid
}
func (r Recruit) Area() string {
	return r.area
}
func (r Recruit) Imgurl() string {
	return r.imgurl
}
func (r Recruit) Man() uint16 {
	return r.man
}
func (r Recruit) Woman() uint16 {
	return r.woman
}
func (r Recruit) VacantMan() uint16 {
	return r.vacant_man
}
func (r Recruit) VacantWoman() uint16 {
	return r.vacant_woman
}
func (r Recruit) Commnt() string {
	return r.commnt
}
func (r Recruit) Date() time.Time {
	return r.date
}
