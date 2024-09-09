package request

type UserDTO struct {
	Name     string `json:"name" example:"Eduardo Melo"`
	Email    string `json:"email" example:"dudu@gmail.com"`
	Password string `json:"password" example:"123456"`
}

type UserLoginDTO struct {
	Email    string `json:"email" example:"dudu@gmail.com"`
	Password string `json:"password" example:"123456"`
}

type UserDeleteDTO struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
