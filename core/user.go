package core

import (
	"context"
	"database/sql"
	"time"

	"github.com/titosunu/wallet-go/payloads"
)

type User struct {
	ID       int64  `db:"id"`
	FullName string `db:"full_name"`
	Phone    string `db:"phone"`
	Email    string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
	EmailVerifiedAtDB sql.NullTime `db:"email_verified"`
	EmailVerifiedAt time.Time `db:"-"`
}

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	Insert(ctx context.Context,user *User) error
	Update(ctx context.Context,user *User) error
}

type UserService interface {
	Authenticate(ctx context.Context, req payloads.AuthReq) (payloads.AuthRes, error)
	ValidateToken(ctx context.Context, token string) (payloads.UserData, error)
	Register(ctx context.Context, req payloads.UserRegisterRequest) (payloads.UserRegisterResponse, error)
	ValidateOTP(ctx context.Context, req payloads.ValidateOtpRequest) error
}