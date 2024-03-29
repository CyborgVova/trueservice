package postgres

import (
	"fmt"
	"os"

	"github.com/cyborgvova/trueservice/internal/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func NewStorage() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func (s *Storage) GetRaw() *storage.UsersTable {
	return &storage.UsersTable{}
}

func (s *Storage) SetRaw(u *storage.UsersTable) error {
	err := storage.ErrRecordExists
	return err
}
