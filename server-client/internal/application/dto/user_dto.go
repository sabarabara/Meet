package dto

import "github.com/google/uuid"

type UserDTO struct {
	userid           *uuid.UUID
	username         string
	imgurl           string
	pronunciation    string
	selfintroduction string
	stars            float32
}

func NewUserDTO(
	userid *uuid.UUID,
	username string,
	imgurl string,
	pronunciation string,
	selfintroduction string,
	stars float32,
) UserDTO {
	return UserDTO{
		userid:           userid,
		username:         username,
		imgurl:           imgurl,
		pronunciation:    pronunciation,
		selfintroduction: selfintroduction,
		stars:            stars,
	}
}

func (u UserDTO) Userid() *uuid.UUID {
	return u.userid
}
func (u UserDTO) Username() string {
	return u.username
}
func (u UserDTO) Imgurl() string {
	return u.imgurl
}
func (u UserDTO) Pronunciation() string {
	return u.pronunciation
}
func (u UserDTO) Selfintroduction() string {
	return u.selfintroduction
}
func (u UserDTO) Stars() float32 {
	return u.stars
}
