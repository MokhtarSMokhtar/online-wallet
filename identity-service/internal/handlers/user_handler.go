package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/jwt"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/models"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/repository"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/shared"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type UserHandler struct {
	Repo *repository.UserRepository
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		Repo: repo,
	}
}

func (uh *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-Type", "application/json")
	var req models.RegisterUserRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			http.Error(w, fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset), http.StatusBadRequest)
		case errors.Is(err, io.EOF):
			http.Error(w, "Request body must not be empty", http.StatusBadRequest)
		case errors.As(err, &unmarshalTypeError):
			http.Error(w, fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset), http.StatusBadRequest)
		default:
			http.Error(w, "Unable to process request", http.StatusBadRequest)
		}
		log.Printf("Error decoding JSON: %v", err)
		return
	}
	//hash the password
	hashedPassword, err := shared.HashPassword(req.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	user := models.User{
		FullName:     req.FullName,
		Email:        strings.ToLower(req.Email),
		Country:      req.Country,
		CountryCode:  req.CountryCode,
		Phone:        req.Phone,
		PasswordHash: hashedPassword,
		UserType:     req.UserType,
		Gender:       req.Gender,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := uh.Repo.AddNew(&user); err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	claims := jwt.Claims{
		UserId: strconv.Itoa(user.Id),
		Email:  user.Email,
		Phone:  user.Phone,
	}
	//create the user token
	token, err := jwt.GenerateToken(claims)
	w.WriteHeader(http.StatusCreated)
	response := models.UserLogin{
		Token:    token,
		Email:    user.Email,
		Phone:    user.Phone,
		FullName: user.FullName,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
