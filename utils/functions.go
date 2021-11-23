package utils

import (
	"AuthenticationService/dal"
	"AuthenticationService/dao"
	"github.com/gin-gonic/gin"
	"time"
)

func GetUserByToken(c *gin.Context) (bool, DAO.User) {
	var user DAO.User
	claim, _ := GetToken(c)
	exist, user := DAL.GetUserByUsername(claim.Username)
	if !exist {
		return false, user
	}
	return true, user
}

func GetCurrentTime() time.Time {
	utc := (3 * time.Hour) + (30 * time.Minute)
	return time.Now().Add(utc)
}
