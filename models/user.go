package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

type UserDTO struct {
	ID       uint   `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

func ToUser(dto UserDTO) User {
	return User{Username: dto.Username, Password: dto.Password, Email: dto.Email}
}

func ToUserDTO(user User) UserDTO {
	return UserDTO{ID: user.ID, Username: user.Username, Password: user.Password, Email: user.Email}
}
