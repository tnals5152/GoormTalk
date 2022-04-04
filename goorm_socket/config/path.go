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
	Path.ProfileImage = filepath.Join(os.Getenv("FILE_PATH"), "profile_image/")
	if _, err := os.Stat(Path.ProfileImage); os.IsNotExist(err) {
		err = os.MkdirAll(Path.ProfileImage, 0775)
		utils.ErrorCheck(err)
	}
}
