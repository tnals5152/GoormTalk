package api

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"goorm_socket/config"
	"goorm_socket/models"
	"goorm_socket/utils"
)

func Login(c *gin.Context) { //로그인 함수
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	utils.ErrorCheck(err)

	var data map[string]interface{}
	json.Unmarshal([]byte(value), &data)
	/*data = {
		"id": "tnals5152@gmail.com",
		"passwd": "password",
	}*/

	// models.LoginCheck(data["id"].(string), data["password"].(string))
	password := sha512.Sum512([]byte(data["passwd"].(string)))
	user := &models.User{
		Username: data["id"].(string),
		Password: string(password[:]),
	}

	fmt.Println(user)
	var userCount int64
	// DB에서 일치하는 유저 검색
	config.GetDB.Model(user).Count(&userCount)

	config.GetDB.Model(user).Where(user).First(user)
	fmt.Println("😋", user)

	if userCount >= 1 { //일치하는 유저 있음
		c.JSON(http.StatusOK, user)
	}
}
