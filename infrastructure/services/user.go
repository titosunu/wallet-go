package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/titosunu/wallet-go/core"
	"github.com/titosunu/wallet-go/infrastructure/utils"
	"github.com/titosunu/wallet-go/payloads"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository  core.UserRepository
	cacheRepository core.CacheRepository
}

func NewUser(userRepository core.UserRepository,
	cacheRepository core.CacheRepository) core.UserService {
	return &userService{
		userRepository:  userRepository,
		cacheRepository: cacheRepository,
	}
}

// Authenticate implements core.UserService.
func (u *userService) Authenticate(ctx context.Context, req payloads.AuthReq) (payloads.AuthRes, error) {
	user, err := u.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return payloads.AuthRes{}, core.ErrUserNotFound
		}
		return payloads.AuthRes{}, err
	}

	if user == (core.User{}) {
		return payloads.AuthRes{}, core.ErrAuthFailed
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return payloads.AuthRes{}, core.ErrAuthFailed
	}

	token := utils.GenerateRandomString(16)

	userJson, _ := json.Marshal(user)
	_ = u.cacheRepository.Set("user:"+token, userJson)

	return payloads.AuthRes{
		Token: token,
	}, nil
}

// ValidateToken implements core.UserService.
func (u *userService) ValidateToken(ctx context.Context, token string) (payloads.UserData, error) {
	data, err := u.cacheRepository.Get("user:" + token)
	if err != nil {
		return payloads.UserData{}, core.ErrAuthFailed
	}
	var user core.User
	_ = json.Unmarshal(data, &user)

	return payloads.UserData{
		ID:       user.ID,
		FullName: user.FullName,
		Phone:    user.Phone,
		Username: user.Username,
	}, nil
}

// Register implements core.UserService.
func (u *userService) Register(ctx context.Context, req payloads.UserRegisterRequest) (payloads.UserRegisterResponse, error) {
	exist, err := u.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return payloads.UserRegisterResponse{}, err
	}

	if exist != (core.User{}) {
		return payloads.UserRegisterResponse{}, core.ErrUsernameTaken
	}
	
	user := core.User{
		FullName: req.FullName,
		Phone: req.Phone,
		Email: req.Email,
		Password: req.Password,
	}

	err = u.userRepository.Insert(ctx, &user)
	if err != nil {
		return payloads.UserRegisterResponse{}, err
	}

	otpCode := utils.GenerateRandomNumber(4)
	referenceId := utils.GenerateRandomString(16)
	//...
}

// ValidateOTP implements core.UserService.
func (u *userService) ValidateOTP(ctx context.Context, req payloads.ValidateOtpRequest) error {
	panic("unimplemented")
}