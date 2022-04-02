package models

import (
	"fmt"
	"goorm_socket/config"
)

//모든 model Migrate 함수
func MigrateAllTable() {
	//delete column은 되지 않음 -> DropColumn이용
	config.SetDB.AutoMigrate(&User{})
	config.SetDB.AutoMigrate(&FriendsRelationship{})
	config.SetDB.AutoMigrate(&Room{})
	config.SetDB.AutoMigrate(&RoomUser{})
	config.SetDB.AutoMigrate(&Message{})
	config.SetDB.AutoMigrate(&File{})
	config.SetDB.AutoMigrate(&Notice{})
	config.SetDB.AutoMigrate(&Link{})

	//FriendsRelationship 테이블과 User테이블 조인한 결과 테스트 코드
	var friends []FriendsRelationship
	result := config.SetDB.Model(&FriendsRelationship{}).Joins("User").Find(&friends)
	fmt.Println(result)

}
