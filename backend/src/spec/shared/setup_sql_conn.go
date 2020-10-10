package testshared

import (
	"os"

	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/infra/mysql"
	"github.com/kelseyhightower/envconfig"
)

// NewTestDb はテストDBのコネクションを返す
func NewTestDb() *mysql.SQLHandler {
	var dbEnv mysql.DbEnv
	envconfig.Process("mysql", &dbEnv)
	dbEnv.Database = os.Getenv("MYSQL_TEST_DATABASSE:")

	return mysql.NewSqlHandler(dbEnv)
}
