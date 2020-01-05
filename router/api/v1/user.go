package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lyhe/auth-micro/client"
	"github.com/lyhe/auth-micro/proto/proto_files/user"
	"github.com/lyhe/auth-micro/util"
	"log"
	"time"
)

var (
	us = client.UserClient()
)

func RegistryUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	nickname := c.PostForm("nickname")
	user := &auth_srv_user.User{
		Username: username,
		Password: password,
		Nickname: nickname,
	}
	resp, err := us.CreateUser(context.Background(), user)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "error",
			"code":    -1,
		})
		return
	}

	result := gin.H{}
	if resp.Success {
		result["message"] = "success"
		result["code"] = 0
	} else {
		result["message"] = resp.Msg
		result["code"] = -1
	}
	c.JSON(200, result)
}

func UserLogin(c *gin.Context) {
	name := c.PostForm("username")
	password := c.PostForm("password")
	resp, err := us.UserLogin(context.Background(), &auth_srv_user.LoginRequest{
		Username: name,
		Password: password,
	})
	if err != nil {
		c.JSON(500, gin.H{"message": "fail"})
		return
	}
	expiration := time.Hour * 24 * 3
	token, err := util.GenerateToken(resp.Nickname, expiration)
	// 设置缓存
	c.JSON(200, gin.H{"token": token})
}
