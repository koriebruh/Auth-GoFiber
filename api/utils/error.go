package utils

import "errors"

var ErrAuthFailed = errors.New("error authentication")

//func GetHttpStatus(err error) int {
//	switch err {
//	case errors.Is(err, ErrAuthFailed):
//		return 401
//	default:
//		return 500
//	}
//}

func GetHttpStatus(err error) int {
	if errors.Is(err, ErrAuthFailed) {
		return 401
	}
	// Tambahkan kondisi lain di sini jika ada error tambahan yang perlu diperiksa
	return 500
}
