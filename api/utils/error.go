package utils

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidCredential = errors.New("invalid credential")

var ErrAuthFailed = errors.New("error authentication")
