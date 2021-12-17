package config

import (
	DAO "AuthenticationService/dao"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(mysql.Open(DbURL(BuildDBConfig())), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(
		&DAO.User{},
	)

}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	port_num, _ := strconv.Atoi(os.Getenv("user_auth_port_var"))
	dbConfig := DBConfig{
		Host:     os.Getenv("user_auth_host_var"),
		Port:     port_num,
		User:     "root",
		Password: os.Getenv("user_auth_pass_var"),
		DBName:   os.Getenv("user_auth_name_var"),
	}
	return &dbConfig
}
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
