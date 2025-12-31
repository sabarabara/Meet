package user

import (
	"errors"

	"github.com/google/uuid"
)

const (
	errorEmptyUserID           = "ユーザーIDが空になっています"
	errorEmptyUsername         = "ユーザー名が空になっています"
	errorEmptyImgurl           = "画像URLが空になっています"
	errorEmptyPronunciation    = "代名詞が空になっています"
	errorEmptySelfintroduction = "自己紹介が空になっています"
	errorInvalidStars          = "評価が無効です"
)

type User struct {
	userid           uuid.UUID
	username         string
	imgurl           string
	pronunciation    string
	selfintroduction string
	stars            float32
}

func NewUser(userid uuid.UUID, username string, imgurl string, pronunciation string, selfintroduction string, stars float32) (User, error) {
	if userid == uuid.Nil {
		return User{}, errors.New(errorEmptyUserID)
	}
	if username == "" {
		return User{}, errors.New(errorEmptyUsername)
	}
	if imgurl == "" {
		return User{}, errors.New(errorEmptyImgurl)
	}
	if pronunciation == "" {
		return User{}, errors.New(errorEmptyPronunciation)
	}
	if selfintroduction == "" {
		return User{}, errors.New(errorEmptySelfintroduction)
	}
	if stars < -100 || stars > 100 {
		return User{}, errors.New(errorInvalidStars)
	}
	return User{
		userid:           userid,
		username:         username,
		imgurl:           imgurl,
		pronunciation:    pronunciation,
		selfintroduction: selfintroduction,
		stars:            stars,
	}, nil
}

func (u User) Userid() uuid.UUID {
	return u.userid
}
func (u User) Username() string {
	return u.username
}
func (u User) Imgurl() string {
	return u.imgurl
}
func (u User) Pronunciation() string {
	return u.pronunciation
}
func (u User) Selfintroduction() string {
	return u.selfintroduction
}
func (u User) Stars() float32 {
	return u.stars
}
