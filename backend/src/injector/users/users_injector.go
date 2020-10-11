package injector

import (
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/infra/mysql"
	handler "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler/auth"
)

func InjectUserHandler(sqlHandler *mysql.SQLHandler) handler.IAuthHandler {

}
