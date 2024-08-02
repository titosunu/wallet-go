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

func (u userRepository) FindByID(ctx context.Context, id int64) (user core.User, err error) {
	query := "SELECT id, full_name, phone, email, username, password FROM users WHERE id = ?"
	row := u.db.QueryRowContext(ctx, query, id)
	err = row.Scan(&user.ID, &user.FullName, &user.Phone, &user.Email, &user.Username, &user.Password)
	return
}

func (u userRepository) FindByUsername(ctx context.Context, username string) (user core.User, err error) {
	query := "SELECT id, full_name, phone, email, username, password FROM users WHERE username = ?"
	row := u.db.QueryRowContext(ctx, query, username)
	err = row.Scan(&user.ID, &user.FullName, &user.Phone, &user.Email, &user.Username, &user.Password)
	return
}
