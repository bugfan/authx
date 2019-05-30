package api

import (
	"github.com/gin-gonic/gin"
)

func NewAPIServer() {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.ErrorLogger())
	return r
}

type Server interface {
	Run()
}
