package apis

import (
	"authx/apis/user"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	Register(g *gin.RouterGroup)
}

var controllers = make([]Controller, 0)

func RegisterController(c Controller) {
	controllers = append(controllers, c)
}

func NewAPIServer() *APIServer {
	s := &APIServer{
		G: gin.Default(),
	}
	s.G.Use(gin.Recovery())
	s.G.Use(gin.ErrorLogger())
	api := s.G.Group("/api")
	api = api.Group("/auth")
	{
		api.POST("/login", user.Login)
		api.POST("/logout", user.Logout)
		api.POST("/logon", user.Logon)
	}
	return s
}

type APIServer struct {
	G *gin.Engine
}
