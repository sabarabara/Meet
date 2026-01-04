package roomrepoimpl

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

func TestInsertRoom_Postgres(t *testing.T) {
	db := setupPostgres(t)

	tx := db.Begin()
	defer tx.Rollback()

	repo := NewRoomRepoImpl(tx)

	recruitID := uuid.MustParse("e2f46b58-a5d2-4376-a2f3-3df243b2c2e5")
	userID := uuid.MustParse("fcf49967-0058-4051-a704-22bd99078606")

	dto := dto.NewRoomDTO(
		recruitID,
		userID,
		"owner",
		false,
	)

	if err := repo.InsertRoom(dto); err != nil {
		t.Fatalf("InsertRoom failed: %v", err)
	}

	var entity RoomEntity
	if err := tx.Table("rooms").
		Where("recruitid = ? AND userid = ?", recruitID, userID).
		First(&entity).Error; err != nil {
		t.Fatalf("failed to fetch record: %v", err)
	}

	if entity.Recruitid != recruitID {
		t.Errorf("expected Recruitid %v, got %v", recruitID, entity.Recruitid)
	}
	if entity.Userid != userID {
		t.Errorf("expected Userid %v, got %v", userID, entity.Userid)
	}
	if entity.Role != "owner" {
		t.Errorf("expected Role %q, got %q", "owner", entity.Role)
	}
	if entity.Isfinish {
		t.Errorf("expected Isfinished false, got true")
	}
}

func TestDeleteRoom_Postgres(t *testing.T) {
	db := setupPostgres(t)
	tx := db.Begin()
	defer tx.Rollback()

	repo := NewRoomRepoImpl(tx)

	recruitID := uuid.MustParse("e2f46b58-a5d2-4376-a2f3-3df243b2c2e5")
	userID := uuid.MustParse("fcf49967-0058-4051-a704-22bd99078606")

	// 先に insert
	dto := dto.NewRoomDTO(
		recruitID,
		userID,
		"member",
		false,
	)

	if err := repo.InsertRoom(dto); err != nil {
		t.Fatalf("InsertRoom failed: %v", err)
	}

	var entity RoomEntity
	if err := tx.Table("rooms").
		Where("recruitid = ? AND userid = ?", recruitID, userID).
		First(&entity).Error; err != nil {
		t.Fatalf("failed to fetch record: %v", err)
	}

	// delete
	if err := repo.DeleteRoom(entity.Roomid); err != nil {
		t.Fatalf("DeleteRoom failed: %v", err)
	}

	// 存在しないことを確認
	err := tx.Table("rooms").
		Where("roomid = ?", entity.Roomid).
		First(&RoomEntity{}).Error

	if err == nil {
		t.Fatalf("expected record to be deleted, but found")
	}
	if err != gorm.ErrRecordNotFound {
		t.Fatalf("unexpected error: %v", err)
	}
}
