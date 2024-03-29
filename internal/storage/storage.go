package storage

import "errors"

var (
	ErrUrlNotFound  = errors.New("url not found")
	ErrUrlExists    = errors.New("url exists")
	ErrRecordExists = errors.New("record exists")
)

type UsersTable struct {
	ID   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(30)"`
	Role string `json:"role" gorm:"type:varchar(150)"`
}

type DBHandler interface {
	GetRaw() *UsersTable
	SetRaw(*UsersTable) error
}
