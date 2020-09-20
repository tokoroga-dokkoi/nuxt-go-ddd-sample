package handler

import handler "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler/auth"

type IAppHandler interface {
	handler.IAuthHandler
}

func NewAppHandler(handler handler.IAuthHandler) IAppHandler {
	return handler
}
