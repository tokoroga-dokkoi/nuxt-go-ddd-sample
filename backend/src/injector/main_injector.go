package injector

import (
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/infra/mysql"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler"
	"github.com/kelseyhightower/envconfig"
)

func InjectDb() *mysql.SQLHandler {
	var dbEnv mysql.DbEnv
	envconfig.Process("mysql", &dbEnv)
	return mysql.NewSqlHandler(dbEnv)
}

func InjectHandlers() handler.IAppHandler {
	sqlHandler := InjectDb()
	// auth injector
	authHandler := InjectAuthHandler(sqlHandler)

	// new app handler
	appHandler := handler.NewAppHandler(authHandler)

	return appHandler
}
