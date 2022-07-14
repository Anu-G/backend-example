package manager

import (
	"wmb-rest-api/auth"
	"wmb-rest-api/delivery/middleware"
)

type MiddlewareManagerInterface interface {
	AuthMiddleware() middleware.AuthTokenMiddlewareInterface
}

type middlewareManager struct {
	auth auth.TokenInterface
}

func NewMiddleware(auth auth.TokenInterface) MiddlewareManagerInterface {
	return &middlewareManager{
		auth: auth,
	}
}

func (mm *middlewareManager) AuthMiddleware() middleware.AuthTokenMiddlewareInterface {
	return middleware.NewTokenValidator(mm.auth)
}
