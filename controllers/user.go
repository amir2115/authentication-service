package DAL

import (
	"AuthenticationService/dal"
	"AuthenticationService/dao"
	"AuthenticationService/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var input DAO.UserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 403, "message": utils.Messages[403], "error": err.Error()})
		return
	}
	exist, user := DAL.GetUserByUsername(input.Username)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 421, "message": utils.Messages[421]})
		return
	}
	authenticated := DAO.CheckPasswordHash(input.Password, user.Password)
	if !authenticated {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 422, "message": utils.Messages[422]})
	}
	token, err := utils.CreateTokens(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 423, "message": utils.Messages[423]})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "token": token, "name": user.Name})
}
