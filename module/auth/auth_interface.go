package auth

import (
	"go_trunk_based/dto"
	authEntity "go_trunk_based/module/auth/entity"
)

type RepositoryInterface interface {
	ForgotPassword(userId uint64, token string) (*authEntity.ForgotPassword, dto.ResponseError)
	ResetPassword(input dto.ResetPassword) dto.ResponseError
}

type UseCaseInterface interface {
	ForgotPassword(input dto.ForgotPassword) dto.ResponseError
	ResetPassword(input dto.ResetPassword) dto.ResponseError
	Login(input dto.LoginRequest) (*dto.LoginResponse, dto.ResponseError)
}
