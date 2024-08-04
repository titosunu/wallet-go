package core

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrAuthFailed = errors.New("authentication failed")
var ErrUsernameTaken = errors.New("username already use")