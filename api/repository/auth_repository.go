package repository

import (
	"context"
	"koriebruh/restful/api/model/domain"
)

type AuthRepository interface {
	Register(ctx context.Context, user *domain.User) error
	UpdateAcc(ctx context.Context, id string, user *domain.User) error
	DeleteAcc(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (domain.User, error)
	FindByUserName(ctx context.Context, username string) (domain.User, error)
}
