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

//curl -d '{"id":"soomin@genielove.com", "password":"passsword"}' -X POST localhost:8000/api/v1/login
// CollectHost godoc
// @Summary Host information collection.
// @Description If it already exists, the changeable information is updated, and in the case of a new host, it is created and returned.
// @Accept json
// @Produce json
// @Param user body RequestData true "User ID and password"
// @Success 200 {object} models.User
// @Failure 400 {object} models.User
// @Failure 404 {object} models.User
// @Failure 500 {object} models.User
// @Router /login [post]
func Login(c *gin.Context) { //로그인 함수
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	utils.ErrorCheck(err)

	var data map[string]interface{}
	json.Unmarshal(value, &data)
	fmt.Println("😊", data)
	/*data = {
		"id": "tnals5152@gmail.com",
		"password": "password",
	}*/

	// models.LoginCheck(data["id"].(string), data["password"].(string))
	password := sha512.Sum512([]byte(data["password"].(string)))

	var users []models.User
	user := &models.User{
		Username: data["id"].(string),
		Password: string(password[:]),
	}

	// var userCount int64
	// DB에서 일치하는 유저 검색
	// config.GetDB.Model(user).Where(user).Count(&userCount)

	// config.GetDB.Model(user).Where(user).Update("password", string(password[:]))
	// user.Password = string(password[:])
	// config.SetDB.Save(user)
	//result.RowsAffected - result개수
	result := config.GetDB.Model(user).Where(user).Find(&users)
	fmt.Println("😀", users)

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
	/*data = {
		"id": "tnals5152@gmail.com",
		"password": "password",
		"name": "지수민",
		"profile": file or nil,
	}*/
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	utils.ErrorCheck(err)

	var user models.User
	json.Unmarshal(value, &user)

	profileImage, err := c.FormFile("profile_image")
	utils.ErrorCheck(err)

	if err == nil {
		user.ProfileImage = fmt.Sprintf("%s/%s/%s",
			config.Path.ProfileImage, user.Username, profileImage.Filename)
		err = c.SaveUploadedFile(profileImage, user.ProfileImage)
		utils.ErrorCheck(err)
	}

	result := config.SetDB.Model(&user).Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"user": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}

}
