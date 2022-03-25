package api

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"goorm_socket/config"
	"goorm_socket/models"
	"goorm_socket/utils"
)

//curl -d '{"id":"soomin@genielove.com", "passwd":"passsword"}' -X POST localhost:8000/login
func Login(c *gin.Context) { //로그인 함수
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	utils.ErrorCheck(err)

	var data map[string]interface{}
	json.Unmarshal(value, &data)
	/*data = {
		"id": "tnals5152@gmail.com",
		"password": "password",
	}*/

	// models.LoginCheck(data["id"].(string), data["password"].(string))
	password := sha512.Sum512([]byte(data["passwd"].(string)))

	var users []models.User
	user := &models.User{
		Username: data["id"].(string),
		Password: string(password[:]),
	}

	// var userCount int64
	// DB에서 일치하는 유저 검색
	// config.GetDB.Model(user).Where(user).Count(&userCount)

	//result.RowsAffected - result개수
	result := config.GetDB.Model(user).Where(user).Find(&users)

	if result.RowsAffected == 1 { //일치하는 유저 있음 1개 -> 로그인 성공
		c.JSON(http.StatusOK, gin.H{
			"user": users[0],
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"user": nil,
		})
	}
}

func CreateUser(c *gin.Context) { //회원가입
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	utils.ErrorCheck(err)

	var user models.User
	json.Unmarshal(value, &user)

	profileImage, err = c.FormFile("profile_image")
	utils.ErrorCheck(err)
	if err == nil {
		err = c.SaveUploadFile(profileImage, fmt.Sprintf("%s/profile_image/%s", os.Getenv("FILE_PATH")))
	}

	/*data = {
		"id": "tnals5152@gmail.com",
		"password": "password",
		"name": "지수민",
		"profile": file or nil,
	}*/

}
