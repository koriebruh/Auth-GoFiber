package service

import (
	"context"
	"koriebruh/restful/api/model/web"
)

type AuthService interface {
	Authentication(ctx context.Context, request web.AuthRequest) (web.AuthResponse, error)
	Validate(ctx context.Context, token string) (web.UserData, error)
}
