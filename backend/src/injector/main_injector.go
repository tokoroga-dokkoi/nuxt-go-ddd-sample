package injector

import (
	usersinjector "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/injector/users"

	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/infra/mysql"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler"
	"github.com/kelseyhightower/envconfig"
)

func InjectDb() *mysql.SQLHandler {
	var dbEnv mysql.DbEnv
	envconfig.Process("mysql", &dbEnv)
	return mysql.NewSqlHandler(dbEnv)
}

func InjectHandlers() handler.AppHandler {
	sqlHandler := InjectDb()
	// new user handler
	userHandler := usersinjector.InjectUserHandler(sqlHandler)
	// new app handler
	appHandler := handler.AppHandler{
		UsersHandler: userHandler,
	}

	return appHandler
}
