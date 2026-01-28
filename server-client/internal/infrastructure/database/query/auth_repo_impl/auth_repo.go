package authrepoimpl

import (
	"database/sql"
	"server-client/internal/domain/repository/query"

	"github.com/google/uuid"
)

var _ query.AuthRepo = (*AuthRepoImpl)(nil)

type SessionEntity struct {
	UserID   uuid.UUID `db:"userid"`
	Username string    `db:"username"`
}

type AuthRepoImpl struct {
	DB *sql.DB
}

func NewAuthRepoImpl(db *sql.DB) *AuthRepoImpl {
	return &AuthRepoImpl{
		DB: db,
	}
}

func (a *AuthRepoImpl) GetAuthenticatedUserInfo(sub string, provider string) (string, string, error) {
	query := `
        SELECT 
            u.userid, 
            u.username 
        FROM authentication a
        INNER JOIN users u ON a.userid = u.userid
        WHERE a.sub = $1 AND a.provider = $2
    `

	var entity SessionEntity
	err := a.DB.QueryRow(query, sub, provider).Scan(&entity.UserID, &entity.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", nil
		}
		return "", "", err
	}

	return entity.UserID.String(), entity.Username, nil
}
