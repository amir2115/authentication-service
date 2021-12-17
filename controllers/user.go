package controllers

import (
	"AuthenticationService/dal"
	"AuthenticationService/dao"
	"AuthenticationService/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(c *gin.Context) {
	var input DAO.UserAuthenticate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 403, "message": utils.Messages[403], "error": err.Error()})
		return
	}
	exist, _ := GetUserByToken(input.Token)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 403, "message": utils.Messages[403]})
		return
	}
	c.JSON(http.StatusOK, gin.H{"messageCode": 409, "message": utils.Messages[409]})
}

func Login(c *gin.Context) {
	var input DAO.UserLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 403, "message": utils.Messages[403], "error": err.Error()})
		return
	}
	exist, user := DAL.GetUserByUsername(input.Username)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 407, "message": utils.Messages[407]})
		return
	}
	authenticated := DAO.CheckPasswordHash(input.Password, user.Password)
	if !authenticated {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 408, "message": utils.Messages[408]})
	}
	token, err := utils.CreateTokens(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 406, "message": utils.Messages[406]})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "name": user.Name, "username": user.Username})
}

func Signup(c *gin.Context) {
	var input DAO.UserSignup
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 403, "message": utils.Messages[403], "error": err.Error()})
		return
	}
	exist, _ := DAL.GetUserByUsername(input.Username)
	if exist {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 404, "message": utils.Messages[404]})
		return
	}
	if len(input.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 405, "message": utils.Messages[405]})
		return
	}
	user := DAL.CreateUser(input)
	token, err := utils.CreateTokens(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 406, "message": utils.Messages[406]})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "name": user.Name, "username": user.Username})
}
