package core

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrAuthFailed = errors.New("authentication failed")