package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sharantechuser/go-jwt-security/pkg/handler"
)

func New() *gin.Engine {
	r := gin.Default()

	r.POST("/login", handler.LoginHandler)

	r.GET("/authorise", handler.AuthHandler)

	return r
}
