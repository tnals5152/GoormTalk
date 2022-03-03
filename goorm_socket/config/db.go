package config

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"goorm_socket/utils"
)

var GetDB *gorm.DB
var SetDB *gorm.DB

type User struct {
	// gorm.Model
	User_ID      uint64 `gorm:"primary_key"`
	Username     string
	Name         string
	ProfileImage string
}

//각 서버에서 실행 시 디비 연결
func ConnectDB() {
	GetDB, err := gorm.Open(os.Getenv("DB_TYPE"),
		//사용자:비밀번호@tcp(ipAddress:port)
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"),
			os.Getenv("GET_DB_HOST"), os.Getenv("GET_DB_PORT"),
			os.Getenv("DB_NAME")),
	)
	utils.ErrorCheck(err)
	fmt.Println(GetDB)

	SetDB, err := gorm.Open(os.Getenv("DB_TYPE"),
		//사용자:비밀번호@tcp(ipAddress:port)
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"),
			os.Getenv("SET_DB_HOST"), os.Getenv("SET_DB_PORT"),
			os.Getenv("DB_NAME")),
	)
	utils.ErrorCheck(err)
	fmt.Println(SetDB)

	// var user User
	// GetDB.First(&user, 1)

}
