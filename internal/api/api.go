package api

import (
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRoutes() routes {
	r := routes{
		router: gin.Default(),
	}

	v1 := r.router.Group("/api/v1")

	r.addPing(v1)
	r.addUsers(v1)

	return r
}

func (r routes) Run(addr ...string) error {
	return r.router.Run()
}
