package authrepoimpl

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

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

func TestGetAuthenticatedUserInfo_Postgres(t *testing.T) {
	db := setupSQLDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("db.Close error: %v", err)
		}
	}()

	// データが実際に存在するか確認
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM authentication").Scan(&count)
	if err != nil {
		t.Fatalf("failed to count authentication: %v", err)
	}
	log.Printf("Authentication table count: %d", count)

	var authData struct {
		UserID   string
		Sub      string
		Provider string
	}
	err = db.QueryRow("SELECT userid::text, sub, provider FROM authentication LIMIT 1").Scan(&authData.UserID, &authData.Sub, &authData.Provider)
	if err != nil {
		t.Logf("failed to fetch auth data: %v", err)
	} else {
		log.Printf("Auth data: UserID=%s, Sub=%s, Provider=%s", authData.UserID, authData.Sub, authData.Provider)
	}

	repo := NewAuthRepoImpl(db)
	userID, username, err := repo.GetAuthenticatedUserInfo("test_user_sub_001", "cognito")
	if err != nil {
		t.Fatalf("GetAuthenticatedUserInfo failed: %v", err)
	}
	if userID == "" || username == "" {
		t.Fatalf("expected user ID and username, got empty strings")
	}
}
