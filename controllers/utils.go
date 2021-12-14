package controllers

import (
	DAL "AuthenticationService/dal"
	DAO "AuthenticationService/dao"
	"AuthenticationService/utils"
	"github.com/gin-gonic/gin"
)

func GetUserByToken(c *gin.Context) (bool, DAO.User) {
	var user DAO.User
	claim, _ := utils.GetToken(c)
	exist, user := DAL.GetUserByUsername(claim.Username)
	if !exist {
		return false, user
	}
	return true, user
}
