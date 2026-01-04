package reviewrepoimpl

import (
	"fmt"
	"os"
	"testing"

	"server-client/internal/application/dto"

	"github.com/google/uuid"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func setupPostgres(t *testing.T) *gorm.DB {
	t.Helper()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		t.Fatal("DB 環境変数が設定されていません")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable application_name=%s",
		host, port, user, password, dbname, "ess-server-test",
	)

	db, err := gorm.Open(driver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		t.Fatalf("failed to connect db: %v", err)
	}

	return db
}

func TestInsertReview_Postgres(t *testing.T) {
	db := setupPostgres(t)

	tx := db.Begin()
	defer tx.Rollback()

	repo := NewReviewRepoImpl(tx)

	roomID := uuid.MustParse("746c5625-d9c1-4fad-bbe5-f8026a0482a1")
	writerID := uuid.MustParse("0a375ec0-d459-42f2-8160-f805baa0c163")
	evaluatedUserID := uuid.MustParse("2b34f2b5-799c-4046-83b8-96f385a5ac24")
	comment := "とても良い部屋でした"

	dto := dto.NewReviewDTO(
		nil,
		roomID,
		writerID,
		evaluatedUserID,
		comment,
	)

	if err := repo.InsertReview(dto); err != nil {
		t.Fatalf("InsertReview failed: %v", err)
	}

	var entity ReviewEntity
	if err := tx.Table("review").
		Where("roomid = ?", roomID).
		First(&entity).Error; err != nil {
		t.Fatalf("failed to fetch record: %v", err)
	}

	if entity.Roomid != roomID {
		t.Errorf("expected RoomID %v, got %v", roomID, entity.Roomid)
	}
	if entity.Writer != writerID {
		t.Errorf("expected WriterID %v, got %v", writerID, entity.Writer)
	}
	if entity.EvaluatedUser != evaluatedUserID {
		t.Errorf("expected EvaluatedUserID %v, got %v", evaluatedUserID, entity.EvaluatedUser)
	}
	if entity.Comment != comment {
		t.Errorf("expected Comment %q, got %q", comment, entity.Comment)
	}
}
