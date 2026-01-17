package common

type UserDTO struct {
	userID        string
	username      string
	stars         float32
	imgURL        string
	pronunciation string
	selfIntroduce string
}

func NewUserDTO(userID, username string, stars float32, imgURL, pronunciation, selfIntroduce string) *UserDTO {
	return &UserDTO{
		userID:        userID,
		username:      username,
		stars:         stars,
		imgURL:        imgURL,
		pronunciation: pronunciation,
		selfIntroduce: selfIntroduce,
	}
}

func (u *UserDTO) UserID() string {
	return u.userID
}

func (u *UserDTO) Username() string {
	return u.username
}

func (u *UserDTO) Stars() float32 {
	return u.stars
}

func (u *UserDTO) ImgURL() string {
	return u.imgURL
}

func (u *UserDTO) Pronunciation() string {
	return u.pronunciation
}

func (u *UserDTO) SelfIntroduce() string {
	return u.selfIntroduce
}
