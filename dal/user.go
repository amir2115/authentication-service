package DAL

import (
	"AuthenticationService/config"
	"AuthenticationService/dao"
	"AuthenticationService/utils"
)

func CreateUser(username, password, name string) DAO.User {
	hash, err := DAO.HashPassword(password)
	if err != nil {
		return DAO.User{}
	}
	user := DAO.User{
		Username: username,
		Password: hash,
		Name:     name,
		Time:     utils.GetCurrentTime(),
	}
	config.DB.Create(&user)
	return user
}

func ChangePassword(user DAO.User, password string) bool {
	hash, err := DAO.HashPassword(password)
	if err != nil {
		return false
	}
	config.DB.Model(&user).Update(DAO.User{Password: hash})
	return true
}

func DeleteUser(id uint) bool {
	var user DAO.User
	if err := config.DB.Where(&DAO.User{ID: id}).Delete(&user).Error; err != nil {
		return false
	}
	return true
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

func GetAllUsers(page uint) ([]DAO.User, error) {
	var users []DAO.User
	var user DAO.User
	offset := (page - 1) * utils.PaginationLimit
	queryBuilder := config.DB.Limit(utils.PaginationLimit).Offset(offset)
	result := queryBuilder.Model(&DAO.User{}).Where(user).Find(&users)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return users, nil
}
