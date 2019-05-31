package user

import (
	"authx/models"
	"authx/utils"
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
	if user.Name == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, map[string]string{"msg": "帐号或者密码不能为空"})
		return
	}
	if !user.Exist() {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	attr := map[string]string{"Username": user.Name}
	// 1.gen jwt
	// jwtStr, _ := utils.GenJWT(attr)
	// 2.save jwt
	// user.JWT = jwtStr
	// 3.setcookie
	utils.SetJWTDataInCookie(ctx.Writer, attr)

	ctx.JSON(http.StatusOK, user.Name)
}

func Logon(ctx *gin.Context) {
	user := models.User{}
	err := ctx.Bind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	if user.Name == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, map[string]string{"msg": "帐号或者密码不能为空"})
		return
	}
	if has, _ := models.GetEngine().Table("user").Where("name=?", user.Name).Exist(); has {
		ctx.JSON(http.StatusBadRequest, map[string]string{"msg": "此帐号已经注册"})
		return
	}
	user.Password = utils.EncryptedPassword(user.Password)
	_, err = models.GetEngine().Insert(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	attr := map[string]string{"Username": user.Name}
	// jwtStr, _ := utils.GenJWT(attr)
	// user.JWT = jwtStr
	utils.SetJWTDataInCookie(ctx.Writer, attr)
	ctx.JSON(http.StatusOK, nil)
}

func Logout(ctx *gin.Context) {
	utils.DeleteJWTCookie(ctx.Writer)
	ctx.JSON(http.StatusOK, nil)
}
