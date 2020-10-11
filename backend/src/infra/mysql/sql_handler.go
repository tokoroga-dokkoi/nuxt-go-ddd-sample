package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// OS Env for DB
type DbEnv struct {
	User     string
	Password string
	Database string
	Host     string
	Port     int
}

// Model
type SQLHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler(dbEnv DbEnv) *SQLHandler {
	mysqlHost := fmt.Sprintf("tcp(%s:%d)", dbEnv.Host, dbEnv.Port)
	connect := dbEnv.User + ":" + dbEnv.Password + "@" + mysqlHost + "/" + dbEnv.Database + "?parseTime=true"
	// Connect DB
	db, err := gorm.Open("mysql", connect)
	if err != nil {
		panic(err.Error())
	}

	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = db
	// Enable Log Output
	sqlHandler.Conn.LogMode(true)
	return sqlHandler
}

func (sqlHandler *SQLHandler) Close() {
	sqlHandler.Conn.Close()
}
