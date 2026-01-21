package userrepoimpl

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"server-client/internal/application/dto"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func setupSQLDB(t *testing.T) *sql.DB {
	t.Helper()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		t.Fatal("DB environment variables are not set")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	if err := db.Ping(); err != nil {
		err := db.Close()
		if err != nil {
			log.Printf("db.Close error: %v", err)
		}
		t.Fatalf("failed to ping db: %v", err)
	}
	return db
}

func TestGetUsers_Postgres(t *testing.T) {
	userid := uuid.MustParse("fcf49967-0058-4051-a704-22bd99078606")
	db := setupSQLDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("db.Close error: %v", err)
		}
	}()

	repo := NewUserRepoImpl(db)
	user, err := repo.GetUserByID(userid)
	if err != nil {
		t.Fatalf("GetUsers failed: %v", err)
	}
	if user == (dto.UserDTO{}) {
		t.Fatalf("expected user, got nil")
	}
}
