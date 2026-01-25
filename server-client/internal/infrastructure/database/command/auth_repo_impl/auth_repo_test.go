package authrepoimpl

import (
	"fmt"
	"log"
	"os"
	"testing"

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

func TestCreateAuthenticatedTable_Postgres(t *testing.T) {
	db := setupPostgres(t)

	tx := db.Begin()
	defer tx.Rollback()

	repo := NewAuthRepoImpl(tx)

	userid := uuid.MustParse("fcf49967-0058-4051-a704-22bd99078606")
	sub := "test_user_sub_001"
	provider := "cognito"

	err := repo.CreateAuthenticatedTable(userid, sub, provider)
	if err != nil {
		t.Fatalf("CreateAuthenticatedTable failed: %v", err)
	}

	log.Printf("Created authentication for UserID: %v", userid)

	var entity AuthEntity
	if err := tx.Table("authentication").
		Where("userid = ?", userid).
		First(&entity).Error; err != nil {
		t.Fatalf("failed to fetch authentication: %v", err)
	}

	if entity.UserID != userid {
		t.Errorf("expected UserID %v, got %v", userid, entity.UserID)
	}
	if entity.Sub != sub {
		t.Errorf("expected Sub %q, got %q", sub, entity.Sub)
	}
	if entity.Provider != provider {
		t.Errorf("expected Provider %q, got %q", provider, entity.Provider)
	}
}
