package recruitrepoimpl

import (
	"fmt"
	"os"
	"testing"
	"time"

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

func TestInsertRecruit_Postgres(t *testing.T) {
	db := setupPostgres(t)

	tx := db.Begin()
	defer tx.Rollback()

	repo := NewRecruitRepoImpl(tx)

	userID := uuid.MustParse("fcf49967-0058-4051-a704-22bd99078606")
	now := time.Now().UTC().Truncate(time.Microsecond)

	dto := dto.NewRecruitDTO(
		userID,
		"Tokyo",
		"https://example.com/image.png",
		2,
		1,
		1,
		0,
		"test comment",
		now,
	)

	if err := repo.InsertRecruit(dto); err != nil {
		t.Fatalf("InsertRecruit failed: %v", err)
	}

	var entity RecruitEntity
	if err := tx.Table("recruit").
		Where("userid = ?", userID).
		First(&entity).Error; err != nil {
		t.Fatalf("failed to fetch record: %v", err)
	}

	if entity.Userid != userID {
		t.Errorf("expected Userid %v, got %v", userID, entity.Userid)
	}
	if entity.Area != "Tokyo" {
		t.Errorf("expected Area %q, got %q", "Tokyo", entity.Area)
	}
	if entity.Imgurl != "https://example.com/image.png" {
		t.Errorf("expected Imgurl %q, got %q", "https://example.com/image.png", entity.Imgurl)
	}
	if entity.Man != 2 {
		t.Errorf("expected Man %d, got %d", 2, entity.Man)
	}
	if entity.Woman != 1 {
		t.Errorf("expected Woman %d, got %d", 1, entity.Woman)
	}
	if entity.Vacant_man != 1 {
		t.Errorf("expected Vacant_man %d, got %d", 1, entity.Vacant_man)
	}
	if entity.Vacant_woman != 0 {
		t.Errorf("expected Vacant_woman %d, got %d", 0, entity.Vacant_woman)
	}
	if entity.Comment != "test comment" {
		t.Errorf("expected Comment %q, got %q", "test comment", entity.Comment)
	}
	if !entity.Date.Equal(now.UTC()) {
		t.Errorf("expected Date %v, got %v", now.UTC(), entity.Date)
	}
}
