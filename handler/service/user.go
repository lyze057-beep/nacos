package service

import (
	"5/exam/nacos/handler/dao"
	"5/exam/nacos/handler/model"
	"5/exam/nacos/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	type RegisterReq struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
		Phone    string `form:"phone" binding:"required"`
	}
	var req RegisterReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logindao := &dao.Register{}
	info, err := logindao.GetUserByDao(req.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "数据库查询失败"})
		return
	}
	if info != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户已存在"})
		return
	}
	password := pkg.MD5(req.Password)
	newuser := model.Users{
		Username: req.Username,
		Password: password,
		Phone:    req.Phone,
	}
	err = logindao.CreateDao(&newuser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "注册失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功", "data": newuser})
}
func Login(c *gin.Context) {
	type LoginReq struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
		Phone    string `form:"phone" binding:"required"`
	}
	var req LoginReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logindao := dao.Register{}
	users, err := logindao.GetUserByDao(req.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "数据库查询失败"})
		return
	}
	if users == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户已存在"})
		return
	}
	if users.Password != pkg.MD5(req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码错误"})
		return
	}
	var user model.Users
	token, _ := pkg.GetJwtToken(int(user.ID))
	c.JSON(http.StatusOK, gin.H{"msg": "登录成功", "data": users, "token": token})
}
