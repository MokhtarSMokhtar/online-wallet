package repository

import (
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/models"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/sql"
)

type UserRepository struct {
	Db *sql.Db
}

func (r *UserRepository) AddNew(u *models.User) error {
	query := "INSERT INTO users (full_name, email, country, country_code, phone, password_hash, salt, user_type, gender, created_at, updated_at)VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11)RETURNING id"

	err := r.Db.Sql.QueryRow(query,
		u.FullName,
		u.Email,
		u.Country,
		u.CountryCode,
		u.Phone,
		u.PasswordHash,
		u.Salt,
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

func (r *UserRepository) GetUserByPhone(phone, countryCode string) (*models.User, error) {
	query := "SELECT id, full_name, email, country, country_code, phone, user_type, gender, password_hash ,salt FROM users WHERE phone = $1 AND country_code = $2"
	u := &models.User{}
	err := r.Db.Sql.QueryRow(query, phone, countryCode).Scan(&u.Id, &u.FullName, &u.Email, &u.Country, &u.CountryCode, &u.Phone, &u.UserType, &u.Gender, &u.PasswordHash, &u.Salt)
	if err != nil {
		return nil, err
	}
	return u, nil

}
