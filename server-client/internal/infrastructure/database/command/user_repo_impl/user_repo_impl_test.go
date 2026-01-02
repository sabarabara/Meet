package userrepoimpl

import (
	"fmt"
	"log"
	"os"
	"testing"

	"server-client/internal/application/dto"

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

	log.Println("DB_HOST:", host)
	log.Println("DB_PORT:", port)
	log.Println("DB_USER:", user)
	log.Println("DB_PASSWORD:", password)
	log.Println("DB_NAME:", dbname)

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

func TestUpsertUser_Insert(t *testing.T) {
	db := setupPostgres(t)

	tx := db.Begin()
	defer tx.Rollback()

	repo := NewUserRepoImpl(tx)

	in := dto.NewUserDTO(
		nil,
		"test-user",
		"https://example.com/icon.png",
		"they/them",
		"hello world",
		4.5,
	)

	res, err := repo.UpsertUser(in)
	if err != nil {
		t.Fatalf("UpsertUser failed: %v", err)
	}

	log.Printf("Inserted UserID: %v", res.Userid())

	var entity UserEntity
	if err := tx.Table("users").
		Where("userid = ?", res.Userid()).
		First(&entity).Error; err != nil {
		t.Fatalf("failed to fetch user: %v", err)
	}

	if entity.Userid != *res.Userid() {
		t.Errorf("expected UserID %v, got %v", res.Userid(), entity.Userid)
	}
	if entity.Username != "test-user" {
		t.Errorf("expected Username %q, got %q", "test-user", entity.Username)
	}
	if entity.Stars != 4.5 {
		t.Errorf("expected Stars %v, got %v", 4.5, entity.Stars)
	}
}

func TestUpsertUser_Update(t *testing.T) {
	db := setupPostgres(t)

	tx := db.Begin()
	defer tx.Rollback()

	repo := NewUserRepoImpl(tx)

	// insert
	in1 := dto.NewUserDTO(
		nil,
		"old-name",
		"old.png",
		"he/him",
		"old intro",
		3.0,
	)

	res1, err := repo.UpsertUser(in1)
	if err != nil {
		t.Fatalf("UpsertUser(insert) failed: %v", err)
	}

	uid := res1.Userid()

	// update
	in2 := dto.NewUserDTO(
		uid,
		"new-name",
		"new.png",
		"she/her",
		"new intro",
		4.8,
	)

	if _, err := repo.UpsertUser(in2); err != nil {
		t.Fatalf("UpsertUser(update) failed: %v", err)
	}

	var entity UserEntity
	if err := tx.Table("users").
		Where("userid = ?", uid).
		First(&entity).Error; err != nil {
		t.Fatalf("failed to fetch user: %v", err)
	}

	if entity.Username != "new-name" {
		t.Errorf("expected updated Username, got %q", entity.Username)
	}
	if entity.Stars != 4.8 {
		t.Errorf("expected updated Stars, got %v", entity.Stars)
	}
}

func TestDeleteUser_Postgres(t *testing.T) {
	db := setupPostgres(t)

	tx := db.Begin()
	defer tx.Rollback()

	repo := NewUserRepoImpl(tx)

	in := dto.NewUserDTO(
		nil,
		"delete-user",
		"img.png",
		"they/them",
		"bye",
		2.0,
	)

	res, err := repo.UpsertUser(in)
	if err != nil {
		t.Fatalf("UpsertUser failed: %v", err)
	}

	if err := repo.DeleteUser(*res.Userid()); err != nil {
		t.Fatalf("DeleteUser failed: %v", err)
	}

	err = tx.Table("users").
		Where("userid = ?", *res.Userid()).
		First(&UserEntity{}).Error

	if err == nil {
		t.Fatalf("expected user to be deleted")
	}
	if err != gorm.ErrRecordNotFound {
		t.Fatalf("unexpected error: %v", err)
	}
}
