package testshared

import (
	"os"

	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/infra/mysql"
	"github.com/kelseyhightower/envconfig"
)

func truncateTables(sqlConn *mysql.SQLHandler) {
	rows, err := sqlConn.Conn.Raw("show tables").Rows()

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			panic(err.Error())
		}
		// truncate table
		if table != "gorp_migrations" {
			sqlConn.Conn.Exec("TRUNCATE TABLE " + table)
		}
	}
}

// NewTestDb はテストDBのコネクションを返す
func NewTestDb() *mysql.SQLHandler {
	var dbEnv mysql.DbEnv
	envconfig.Process("mysql", &dbEnv)
	dbEnv.Database = os.Getenv("MYSQL_TEST_DATABASE")
	return mysql.NewSqlHandler(dbEnv)
}

// CloseTestDbConn はテストDBのコネクションをクローズする
func CloseTestDbConn(sqlConn *mysql.SQLHandler) {
	truncateTables(sqlConn)
	sqlConn.Close()
}
