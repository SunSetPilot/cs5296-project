package dal

import (
	"database/sql"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mock sqlmock.Sqlmock

func InitMockDB() error {
	var (
		db  *sql.DB
		err error
	)
	db, mock, err = sqlmock.New()
	if err != nil {
		return fmt.Errorf("sqlmock open connection error: %v", err)
	}
	DB = new(_DB)
	DB.db, err = gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("gorm open mock connection error: %v", err)
	}
	return nil
}
