package DAL

import (
	"AuthenticationService/config"
	"AuthenticationService/dao"
	"AuthenticationService/utils"
)

func CreateUser(userSignup DAO.UserSignup) DAO.User {
	hash, err := DAO.HashPassword(userSignup.Password)
	if err != nil {
		return DAO.User{}
	}
	user := DAO.User{
		Username: userSignup.Username,
		Password: hash,
		Name:     userSignup.Name,
		Time:     utils.GetCurrentTime(),
	}
	config.DB.Create(&user)
	return user
}

func GetUser(id uint) (bool, DAO.User) {
	var user DAO.User
	if err := config.DB.Where(&DAO.User{ID: id}).First(&user).Error; err != nil {
		return false, user
	}
	return true, user
}

func GetUserByUsername(username string) (bool, DAO.User) {
	var user DAO.User
	if err := config.DB.Where(&DAO.User{Username: username}).First(&user).Error; err != nil {
		return false, user
	}
	return true, user
}
