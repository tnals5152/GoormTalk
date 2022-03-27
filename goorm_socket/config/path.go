package config

type PathStruct struct {
	ProfileImage string
}

var Path *PathStruct

func InitPath() {
	// Path.ProfileImage = fmt.Sprintf("%s/profile_image", os.Getenv("FILE_PATH"))
}
