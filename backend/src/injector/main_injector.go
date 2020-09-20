package injector

import (
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/infra/mysql"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler"
)

func InjectDb() *mysql.SQLHandler {
	return mysql.NewSqlHandler()
}

func InjectHandlers() handler.IAppHandler {
	sqlHandler := InjectDb()
	// auth injector
	authHandler := InjectAuthHandler(sqlHandler)

	// new app handler
	appHandler := handler.NewAppHandler(authHandler)

	return appHandler
}
