package controllers

import (
	DAL "AuthenticationService/dal"
	DAO "AuthenticationService/dao"
	"AuthenticationService/utils"
)

func GetUserByToken(token string) (bool, DAO.User) {
	var user DAO.User
	_, claim := utils.ValidateToken(token)
	exist, user := DAL.GetUserByUsername(claim.Username)
	if !exist {
		return false, user
	}
	return true, user
}
