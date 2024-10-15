package models

type UserRegisteredEvent struct {
	UserID    int    `json:"user_id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Timestamp string `json:"timestamp"`
}