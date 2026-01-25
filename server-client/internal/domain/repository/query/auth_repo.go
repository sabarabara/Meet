package query

type AuthRepo interface {
	GetAuthenticatedUserInfo(sub string, provider string) (userid string, username string, err error)
}
