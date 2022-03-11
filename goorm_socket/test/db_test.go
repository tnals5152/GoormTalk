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
	createRoom(user2)

	fmt.Println("")
}

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
		User:   *user2,
		Friend: *user1,
	}

	config.SetDB.Where(friend).FirstOrCreate(friend).Joins("User")
	fmt.Println(*friend)
}

func createRoom(owner *models.User) {
	room := &models.Room{
		RoomName: "soominWithgenie",
		RoomType: 1,
		Owner:    *owner,
	}

	config.SetDB.Where(room).FirstOrCreate(room).Joins("User")
	fmt.Println(*room)

}
