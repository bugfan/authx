package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func Logon(ctx *gin.Context) {

}

func Logout(ctx *gin.Context) {

}
