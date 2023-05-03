package cmd

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// type Handler struct {
// 	db *gorm.DB
// }

// func NewHandler(db *gorm.DB) *Handler {
// 	return &Handler{db}
// }

func UseDB() (db *gorm.DB, err error) {
	dsn := "sqlserver://sa:Keng1234@localhost?database=TestDB"
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
