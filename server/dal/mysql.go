package dal

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *_DB

type _DB struct {
	db *gorm.DB
}

func (p *_DB) NewRequest(ctx context.Context) *gorm.DB {
	return p.db.Session(&gorm.Session{Context: ctx})
}

func InitDB(dsn string) error {
	var err error
	DB = new(_DB)
	DB.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("gorm open mysql connection error: %v", err)
	}
	return nil
}
