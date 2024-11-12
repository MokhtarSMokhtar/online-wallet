package models

import (
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/enums"
	"time"
)

type User struct {
	Id           int            `json:"id"`
	FullName     string         `json:"full_name"`
	Email        string         `json:"email"`
	Country      string         `json:"country"`
	CountryCode  string         `json:"country_code"`
	Phone        string         `json:"phone"`
	PasswordHash []byte         `json:"password_hash"`
	Salt         []byte         `json:"salt"`
	UserType     enums.UserType `json:"user_type" swaggertype:"string"`
	Gender       string         `json:"gender"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type RegisterUserRequest struct {
	FullName    string         `json:"full_name"`
	Email       string         `json:"email"`
	Country     string         `json:"country"`
	CountryCode string         `json:"country_code"`
	Phone       string         `json:"phone"`
	Password    string         `json:"password"`
	UserType    enums.UserType `json:"user_type" swaggertype:"string"`
	Gender      string         `json:"gender"`
}

type UserLogin struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Token    string `json:"token"`
}

type UserLoginRequest struct {
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
}
