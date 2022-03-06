package config

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"goorm_socket/utils"
)

var GetDB *gorm.DB
var SetDB *gorm.DB

type Test2 struct {
	gorm.Model //id, create_at, update_at, delete_at이 포함되어있는 Model객체
	//id         uint64 `gorm:"colum:id; primary_key"`
	Name       string `gorm:"not null" json:"name"` //소문자 X -> 소문자 사용으로 다른 패키지에서 접근 X -> 생성 X
	address    string
	testString string
}

//각 서버에서 실행 시 디비 연결
func ConnectDB() {
	//사용자:비밀번호@tcp(ipAddress:port)
	dsnGet := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"),
		os.Getenv("GET_DB_HOST"), os.Getenv("GET_DB_PORT"),
		os.Getenv("DB_NAME"))
	GetDB, err := gorm.Open(mysql.Open(dsnGet), &gorm.Config{})
	utils.IfErrorMakePanic(err, "can not connect Get DB")
	fmt.Println(GetDB)

	dsnSet := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"),
		os.Getenv("SET_DB_HOST"), os.Getenv("SET_DB_PORT"),
		os.Getenv("DB_NAME"))
	SetDB, err := gorm.Open(mysql.Open(dsnSet), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
	})
	utils.IfErrorMakePanic(err, "can not connect Set DB")
	migrateAllTable()
	fmt.Println(SetDB)
	// fmt.Println(SetDB.AutoMigrate(Test2{}), "testetset!!😂")
	// db := SetDB.AutoMigrate(Test2{})
	// fmt.Println(db)

}

//모든 model Migrate 함수
func migrateAllTable() {
	SetDB.AutoMigrate(&Test2{})
}
