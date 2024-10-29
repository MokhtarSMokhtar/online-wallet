package models

type UserRegisteredEvent struct {
	UserID    int32  `json:"user_id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Timestamp string `json:"timestamp"`
}
