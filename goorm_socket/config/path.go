package config

import (
	"goorm_socket/utils"
	"os"
	"path/filepath"
)

type PathStruct struct {
	ProfileImage string
}

var Path *PathStruct = &PathStruct{}

func InitPath() {
	// Path.ProfileImage = fmt.Sprintf("%s/profile_image", os.Getenv("FILE_PATH"))
	//mkdirall안 됨 다시 할 것(775안 됨 & FILE_PATH까지만 생성됨)
	Path.ProfileImage = filepath.Join(os.Getenv("FILE_PATH"), "profile_image/")
	if _, err := os.Stat(Path.ProfileImage); os.IsNotExist(err) {
		err = os.MkdirAll(Path.ProfileImage, 775)
		utils.ErrorCheck(err)
	}
}
