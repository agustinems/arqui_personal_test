package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter() {
	Router = gin.Default()
	Router.Use(cors.Default())
}
