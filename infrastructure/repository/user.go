package repository

import (
	"context"
	"database/sql"

	"github.com/titosunu/wallet-go/core"
)

type userRepository struct {
	db *sql.DB
}

func NewUser(con *sql.DB) core.UserRepository {
	return &userRepository{
		db: con,
	}
}

// Insert implements core.UserRepository.
func (u *userRepository) Insert(ctx context.Context, user *core.User) error {
	query := `INSERT INTO users (full_name, username, password, phone, email) VALUES (?, ?, ?, ?, ?) RETURNING id`
	row := u.db.QueryRowContext(ctx, query, user.FullName, user.Username, user.Password, user.Phone, user.Email)
	err := row.Scan(&user.ID)
	return err
}

// Update implements core.UserRepository.
func (u *userRepository) Update(ctx context.Context, user *core.User) error {
	query := `UPDATE users SET full_name = ?, username = ?, password = ?, phone = ?, email = ?, email_verified_at = ? WHERE id = ?`
	_, err := u.db.ExecContext(ctx, query, user.FullName, user.Username, user.Password, user.Phone, user.Email, user.EmailVerifiedAtDB, user.ID)
	return err
}

// Find By Id implements core.UserRepository.
func (u userRepository) FindByID(ctx context.Context, id int64) (user core.User, err error) {
	query := "SELECT id, full_name, phone, email, username, password FROM users WHERE id = ?"
	row := u.db.QueryRowContext(ctx, query, id)
	err = row.Scan(&user.ID, &user.FullName, &user.Phone, &user.Email, &user.Username, &user.Password)
	return
}

// Find By Username implements core.UserRepository.
func (u userRepository) FindByUsername(ctx context.Context, username string) (user core.User, err error) {
	query := "SELECT id, full_name, phone, email, username, password FROM users WHERE username = ?"
	row := u.db.QueryRowContext(ctx, query, username)
	err = row.Scan(&user.ID, &user.FullName, &user.Phone, &user.Email, &user.Username, &user.Password)
	return
}
