package user

import (
	"authx/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	user := models.User{}
	err := ctx.Bind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	if !user.Exist() {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	// 1.gen jwt

	// 2.save jwt
	// 3.setcookie
	ctx.JSON(http.StatusOK, user.Name)
}

func Logon(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
