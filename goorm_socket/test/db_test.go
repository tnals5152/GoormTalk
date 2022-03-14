package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/gorm/clause"

	"goorm_socket/config"
	"goorm_socket/models"
)

//db 테이블에 insert하기
func TestCreateDB(t *testing.T) {
	godotenv.Load("../../.env")
	fmt.Println(os.Getenv("DB_USER"), "🤣")
	config.ConnectDB()
	user1, user2 := createUser()
	createFriendsRelationship(user1, user2)
	room := createRoom(user2)
	createRoomUser(user1, room)
	createRoomUser(user2, room)

	fmt.Println("")
}

//select * from information_schema.table_constraints where table_name = '테이블명';
func createUser() (*models.User, *models.User) {
	user1 := models.User{
		Username: "genie@genielove.com",
		Password: "password",
		Name:     "지니",
	}
	user2 := models.User{
		Username: "soomin@genielove.com",
		Password: "password",
		Name:     "tnals",
	}
	config.SetDB.Where(&user1).FirstOrCreate(&user1)
	config.SetDB.Where(&user2).FirstOrCreate(&user2)
	return &user1, &user2

}

func createFriendsRelationship(user1 *models.User, user2 *models.User) {
	//tnals만 genie를 아는 사이
	friend := &models.FriendsRelationship{
		// User:   *user2,
		// Friend: *user1,
		UserID:   user2.ID,
		FriendID: user1.ID,
	}

	config.SetDB.Where(friend).FirstOrCreate(friend).Joins("User")
	fmt.Println(*friend)
}

func createRoom(owner *models.User) *models.Room {
	room := &models.Room{
		RoomName: "soominWithgenie",
		RoomType: 1,
		UserID:   owner.ID,
	}

	config.SetDB.Where(room).FirstOrCreate(room).Joins("User")
	return room
}

func createRoomUser(user *models.User, room *models.Room) {
	roomUser := &models.RoomUser{}
	var roomUsers []models.RoomUser
	// config.SetDB.Joins("User", "Room").Where(roomUser).FirstOrCreate(roomUser)
	//claus.Associations = 연관된 모든 테이블 데이터 출력
	config.SetDB.Preload(clause.Associations).Where(roomUser).FirstOrCreate(roomUser)
	// config.SetDB.Joins("User", user).Joins("Room", room).Find(&roomUsers)//되는 것
	// config.SetDB.Joins("inner join User on User.id = RoomUser.user_id and User.id=?", 11).Joins("Room", room).Find(&roomUsers)
	// config.SetDB.Joins("User", user).Joins("Room", room).First(&roomUser)
	// config.SetDB.Preload("User", user).Joins("User").FirstOrCreate(roomUser)
	// config.SetDB.Model(user).Preload()
	// config.SetDB.Joins("User", config.SetDB.Model(user).Where(user).First(user)).
	// 	Joins("Room", config.SetDB.Model(room).Where(room)).First(roomUser)
	config.SetDB.Preload(clause.Associations).Where(roomUser).Find(&roomUser)
	fmt.Println("😂", *user)
	for _, r := range roomUsers {
		fmt.Println("👿", r)
	}
	fmt.Println("🤬", roomUser.Room)
	fmt.Println("🤬", roomUser.User)
}
