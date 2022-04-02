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

type loginUser struct {
	Username string `json:"username" example:"soomin@genielove.com"`
	Password string `json:"password" example:"passsword"`
}

//curl -d '{"username":"soomin@genielove.com", "password":"passsword"}' -X POST localhost:8000/api/v1/login
// @Summary login api
// @Description user login
// @Accept json
// @Produce json
// @Param user body loginUser true "User username and password"
// @Success 200 {object} models.User
// @Failure 400 {object} models.User
// @Failure 404 {object} models.User
// @Failure 500 {object} models.User
// @Router /login [post]
func Login(c *gin.Context) { //로그인 함수
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	utils.ErrorCheck(err)

	// var data map[string]interface{}
	var data loginUser
	json.Unmarshal(value, &data)
	password := sha512.Sum512([]byte(data.Password))

	var users []models.User
	user := &models.User{
		Username: data.Username,
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

	if result.RowsAffected == 1 { //일치하는 유저 있음 1개 -> 로그인 성공
		c.JSON(http.StatusOK, gin.H{
			"user": users[0],
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "user",
			"message": "can't find user",
		})
	}
}

// @Summary create user
// @Description create user
// @Accept  multipart/form-data
// @Produce  json
// @Param username formData string true "User email"
// @Param password formData string true "User password"
// @Param name formData string true "User name"
// @Param profile_image formData file true "User profile"
// @Success 200 {object} models.User
// @Failure 400 {object} models.User
// @Failure 404 {object} models.User
// @Failure 500 {object} models.User
// @Router /create-user [post]
func CreateUserAPI(c *gin.Context) { //회원가입
	/*
		"username": "tnals5152@gmail.com",
		"password": "password",
		"name": "지수민",
		"profile": file or nil,
	*/
	var user models.User
	err := c.ShouldBind(&user)
	utils.ErrorCheck(err)
	if !user.CheckIsUnique() {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "user",
			"message": "username is not unique",
		})
		return
	}

	if !PasswordLegnth(user.Password) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "user",
			"message": "password is too short",
		})
		return
	}

	profileImage, err := c.FormFile("profile_image")

	if err == nil {
		//mkdirall 할 것
		user.ProfileImage = fmt.Sprintf("%s/%s/%s",
			config.Path.ProfileImage, user.Username, profileImage.Filename)
		err = c.SaveUploadedFile(profileImage, user.ProfileImage)
		utils.ErrorCheck(err)
	}
	result := user.CreateUser()

	// // c2, err := c.MultipartForm()
	// // fmt.Println(err)
	// // fmt.Println(c2)
	// // utils.ErrorCheck(err)
	// // fmt.Println(user)

	// result := config.SetDB.Model(&user).Create(&user)

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

func PasswordLegnth(password string) bool {
	if len(password) < models.PasswordLength {
		return false
	}
	return true
}
