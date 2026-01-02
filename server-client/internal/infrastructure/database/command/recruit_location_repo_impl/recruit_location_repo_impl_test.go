package recruitlocationrepoimpl

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

func TestInsertRecruitLocation_Postgres(t *testing.T) {
	db := setupPostgres(t)

	tx := db.Begin()
	defer tx.Rollback()

	repo := NewRecruitLocationRepoImpl(tx)

	recruitID := uuid.MustParse("e2f46b58-a5d2-4376-a2f3-3df243b2c2e5")
	latitude := 35.681236
	longitude := 139.767125

	dto := dto.NewRecruitLocationDTO(
		recruitID,
		latitude,
		longitude,
	)

	if err := repo.InsertRecruitLocation(dto); err != nil {
		t.Fatalf("InsertRecruitLocation failed: %v", err)
	}

	var entity RecruitLocationEntity
	err := tx.Table("recruit_location").
		Where("recruitid = ?", recruitID).
		First(&entity).Error
	if err != nil {
		t.Fatalf("failed to fetch record: %v", err)
	}

	if entity.RecruitID != recruitID {
		t.Errorf("expected RecruitID %v, got %v", recruitID, entity.RecruitID)
	}
	if entity.Latitude != latitude {
		t.Errorf("expected Latitude %v, got %v", latitude, entity.Latitude)
	}
	if entity.Longitude != longitude {
		t.Errorf("expected Longitude %v, got %v", longitude, entity.Longitude)
	}
}
