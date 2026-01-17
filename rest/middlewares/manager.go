package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (manager *Manager) Use(middlewares ...Middleware) {
	manager.globalMiddlewares = append(manager.globalMiddlewares, middlewares...)
}

func (manager *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
		handler := next

		for i:=len(middlewares)-1; i>=0; i-- {
			middleware := middlewares[i]
			handler = middleware(handler)
		}
		return handler
}

func (manager *Manager) WrapMux(mux http.Handler) http.Handler {
	handler := mux

	for i:=len(manager.globalMiddlewares)-1; i>=0; i-- {
		middleware := manager.globalMiddlewares[i]
		handler = middleware(handler)
	}
	return handler
}