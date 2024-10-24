package service

import (
	"context"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"koriebruh/restful/api/model/domain"
	"koriebruh/restful/api/model/web"
	"koriebruh/restful/api/repository"
	"koriebruh/restful/api/utils"
	"log/slog"
)

type AuthServiceImpl struct {
	repository      repository.AuthRepository
	cacheRepository utils.CacheRepository
}

func NewAuthService(repo repository.AuthRepository, cacheRepo utils.CacheRepository) AuthService {
	return &AuthServiceImpl{
		repository:      repo,
		cacheRepository: cacheRepo,
	}
}

func (service AuthServiceImpl) Authentication(ctx context.Context, request web.AuthRequest) (web.AuthResponse, error) {
	user, err := service.repository.FindByUserName(ctx, request.UserName)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return web.AuthResponse{}, err
	}

	if user == (domain.User{}) {
		return web.AuthResponse{}, utils.ErrAuthFailed
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return web.AuthResponse{}, utils.ErrAuthFailed
	}

	token := utils.GeneratorRandString(16)
	userJson, _ := json.Marshal(user)
	
	//_ = service.cacheRepository.Set("user"+token, userJson)
	err = service.cacheRepository.Set("user:"+token, userJson)
	if err != nil {
		return web.AuthResponse{}, err
	}

	return web.AuthResponse{
		AccessToken: token,
	}, nil

}

func (service AuthServiceImpl) Validate(ctx context.Context, token string) (web.UserData, error) {
	data, err := service.cacheRepository.Get("user:" + token)
	if err != nil {
		return web.UserData{}, utils.ErrAuthFailed
	}

	var user domain.User
	_ = json.Unmarshal(data, &user)

	return web.UserData{
		ID:       user.ID.String(),
		Name:     user.Name,
		UserName: user.UserName,
		Email:    user.Email,
	}, nil

}
