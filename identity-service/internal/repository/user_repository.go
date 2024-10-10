package repository

import (
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/models"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/sql"
)

type UserRepository struct {
	Db *sql.Db
}

func (r *UserRepository) AddNew(u *models.User) error {
	query := "INSERT INTO users (full_name, email, country, country_code, phone, password_hash, user_type, gender, created_at, updated_at)VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)RETURNING id"

	err := r.Db.Sql.QueryRow(query,
		u.FullName,
		u.Email,
		u.Country,
		u.CountryCode,
		u.Phone,
		u.PasswordHash,
		u.UserType,
		u.Gender,
		u.CreatedAt,
		u.UpdatedAt,
	).Scan(&u.Id)
	if err != nil {
		return err
	}
	return nil
}
