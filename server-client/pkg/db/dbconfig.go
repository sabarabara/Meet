package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	SqlDB  *sql.DB
	GormDB *gorm.DB
)

func InitDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("DB 環境変数が設定されていません")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	SqlDB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("sql.DB 接続失敗:", err)
	}

	GormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: SqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("GORM 接続失敗:", err)
	}

	fmt.Println("SQL & GORM 両方の接続に成功しました")
}
