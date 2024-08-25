package stores

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/asayuki/gopherreads/models"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db}
}

func (s *UserStore) CreateUser(user models.UserAuth) error {
	_, err := s.db.Exec(`
		INSERT INTO users (
			email, password
		) VALUES (?, ?)
	`, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) GetUserByField(field string, value interface{}) (*models.User, error) {
	result := s.db.QueryRow(fmt.Sprintf("SELECT FROM users WHERE %s = ?", field), value)

	user := new(models.User)
	err := scanUser(result, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserStore) CheckAuth(email, password string) (*models.User, error) {
	user, err := s.GetUserByField("email", email)
	if err != nil {
		return nil, err
	}

	if comparePassword(user.Password, []byte(password)) {
		return nil, errors.New("passwords not matching")
	}

	return user, nil
}

func scanUser(row *sql.Row, user *models.User) error {
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	return err
}
