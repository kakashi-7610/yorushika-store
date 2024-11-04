package repository

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

type DatabaseConfig struct {
	Hostname string
	Port     string
	User     string
	Password string
	DbName   string
	Retry    int
}

func NewDatabase(dc *DatabaseConfig) (*Database, error) {
	connectInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dc.User,
		dc.Password,
		dc.Hostname,
		dc.Port,
		dc.DbName,
	)

	log.Println(connectInfo)
	log.Println("Waiting for MySQL container to")
	var err error

	for i := 0; i < dc.Retry; i++ {
		db, err := gorm.Open(mysql.Open(connectInfo), &gorm.Config{})
		if err == nil {
			log.Print("Database connected")
			return &Database{db}, nil
		}

		time.Sleep(2 * time.Second) // 2秒待機
		fmt.Println("Retrying to connect to MySQL...")
	}

	log.Fatalf("failed to connect db. error: %v", err)
	return nil, err
}

func NewDatabaseConfig(hostname string, port string, user string, password string, dbname string, retry int) *DatabaseConfig {
	return &DatabaseConfig{
		Hostname: hostname,
		Port:     port,
		User:     user,
		Password: password,
		DbName:   dbname,
		Retry:    retry,
	}
}

// オートマイグレーション用
// Model更新時のみ使用する
func (db *Database) AutoMigrate() error {
	err := db.db.AutoMigrate(&Product{})
	if err != nil {
		log.Fatalf("AutoMigrate failed. error: %v", err)
		return err
	}

	log.Print("AutoMigrated database")
	return nil
}
