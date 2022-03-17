package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"

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
	room := createRoom(user1)
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
		// RoomName: "hiTest",
		// RoomType: 1,
		// UserID:   owner.ID,
	}
	var users []models.User
	//django orm과 다르게 정방향으로 접근 가능...
	config.SetDB.Joins("User").Last(room)
	config.SetDB.Preload("Room", &models.Room{RoomName: "hiTest"}).Find(&users)
	fmt.Println("🥶 ", users)
	return room
}

func createRoomUser(user *models.User, room *models.Room) {
	roomUser := &models.RoomUser{}
	var roomUsers []models.User
	config.SetDB.Model(user).Association("RoomUser").Find(roomUser)
	// config.SetDB.Preload("RoomUser.Room").Preload("RoomUser", &models.RoomUser{
	// 	Room: *room, User: *user}).Last(&roomUser)
	fmt.Println("😂", *user)
	for _, r := range roomUsers {
		fmt.Println("👿", r.RoomUser, "!!!", r.Room)
	}
	fmt.Println("🤬", roomUser.Room)
	fmt.Println("🤬", roomUser.User)
	fmt.Println("🤬", roomUser)
}
