package recruitrepoimpl

import (
	"database/sql"
	"log"
	"server-client/internal/application/dto"
	"server-client/internal/domain/repository/query"
	"time"

	"github.com/google/uuid"
)

var _ query.RecruitRepo = (*RecruitRepoImpl)(nil)

type recruitEntity struct {
	Recruitid    uuid.UUID `db:"column:recruitid;type:uuid;primaryKey;default:uuid_generate_v4()"`
	Userid       uuid.UUID `db:"column:userid;type:uuid;not null"`
	Area         string    `db:"column:area;type:varchar(255);not null"`
	Imgurl       string    `db:"column:imgurl;type:varchar(255);not null"`
	Man          uint16    `db:"column:man;type:smallint;not null"`
	Woman        uint16    `db:"column:woman;type:smallint;not null"`
	Vacant_man   uint16    `db:"column:vacant_man;type:smallint;not null"`
	Vacant_woman uint16    `db:"column:vacant_woman;type:smallint;not null"`
	Comment      string    `db:"column:comment;type:text;not null"`
	Date         time.Time `db:"column:date;type:timestamp;not null"`
	Latitude     float64   `db:"column:latitude;type:double precision;not null"`
	Longitude    float64   `db:"column:longitude;type:double precision;not null"`
}

type RecruitRepoImpl struct {
	db *sql.DB
}

func NewRecruitRepoImpl(db *sql.DB) *RecruitRepoImpl {
	return &RecruitRepoImpl{db: db}
}

func (r *RecruitRepoImpl) GetRecruits(page int, size int) ([]dto.RecruitDTO, []dto.RecruitLocationDTO, error) {
	offset := (page - 1) * size
	rows, err := r.db.Query(`
    SELECT r.recruitid, r.userid, r.area, r.imgurl, r.man, r.woman,
           r.vacant_man, r.vacant_woman, r.comment, r.date,
           rl.latitude, rl.longitude
    FROM RECRUIT r
    LEFT JOIN RECRUIT_LOCATION rl ON r.recruitid = rl.recruitid
    ORDER BY r.date DESC
    LIMIT $1 OFFSET $2
`, size, offset)

	if err != nil {
		return nil, nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("rows.Close error: %v", err)
			return
		}
	}()

	var recruits []dto.RecruitDTO
	var locations []dto.RecruitLocationDTO
	for rows.Next() {
		var entity recruitEntity
		if err := rows.Scan(&entity.Recruitid, &entity.Userid, &entity.Area, &entity.Imgurl, &entity.Man, &entity.Woman, &entity.Vacant_man, &entity.Vacant_woman, &entity.Comment, &entity.Date, &entity.Latitude, &entity.Longitude); err != nil {
			return nil, nil, err
		}
		recruitDTO := dto.NewRecruitDTO(
			&entity.Recruitid,
			entity.Userid,
			entity.Area,
			entity.Imgurl,
			entity.Man,
			entity.Woman,
			entity.Vacant_man,
			entity.Vacant_woman,
			entity.Comment,
			entity.Date,
		)
		locationDTO := dto.NewRecruitLocationDTO(
			nil,
			entity.Recruitid,
			entity.Latitude,
			entity.Longitude,
		)
		recruits = append(recruits, recruitDTO)
		locations = append(locations, locationDTO)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, err
	}
	return recruits, locations, nil
}
