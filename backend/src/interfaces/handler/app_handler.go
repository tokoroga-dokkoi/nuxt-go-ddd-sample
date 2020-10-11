package handler

import (
	usershandler "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler/users"
)

type AppHandler struct {
	UsersHandler usershandler.UserHandler
}
