package DAO

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Time     time.Time `json:"time"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
