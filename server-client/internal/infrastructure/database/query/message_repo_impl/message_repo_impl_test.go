package messagerepoimpl

import (
	"database/sql"
	"fmt"
	"log"
	"os"
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

func TestGetMessages_Postgres(t *testing.T) {
	db := setupSQLDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("db.Close error: %v", err)
		}
	}()
	roomid := uuid.MustParse("746c5625-d9c1-4fad-bbe5-f8026a0482a1")

	repo := NewMessageRepoImpl(db)
	messages, err := repo.GetMessages(roomid, 1, 1)
	if err != nil {
		t.Fatalf("GetMessages failed: %v", err)
	}
	if messages == nil {
		t.Fatalf("expected messages slice, got nil")
	}
}
