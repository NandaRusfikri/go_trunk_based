package controller

import (
	"github.com/gin-gonic/gin"
	"go_trunk_based/dto"
	"go_trunk_based/module/auth"
	"go_trunk_based/util"
	"net/http"
)

type AuthControllerHTTP struct {
	authUsecase auth.UseCaseInterface
}

func InitAuthControllerHTTP(route *gin.Engine, authUsecase auth.UseCaseInterface) {

	controller := &AuthControllerHTTP{
		authUsecase: authUsecase,
	}
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/auth/login", controller.Login)
	groupRoute.POST("/auth/forgot-password", controller.ForgotPassword)
	groupRoute.POST("/auth/reset-password", controller.ResetPassword)
}

// ResetPassword
//
//	@Tags			Auth
//	@Summary		Reset Password
//	@Description	API for confirm reset password
//	@ID				User-ResetPassword
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			data	body	dto.ResetPassword	true	"body data"
//	@Success		200
//	@Router			/v1/auth/reset-password [post]
func (c *AuthControllerHTTP) ResetPassword(ctx *gin.Context) {

	var input dto.ResetPassword

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, "Request Invalid "+err.Error(), 400, 0, nil)
		return
	}
	err := c.authUsecase.ResetPassword(input)

	if err.Error != nil {
		util.APIResponse(ctx, err.Error.Error(), err.Code, 0, nil)
	} else {
		util.APIResponse(ctx, "Reset Password Success", http.StatusOK, 0, nil)
	}
}

// ForgotPassword
//
//	@Tags			Auth
//	@Summary		Forgot Password
//	@Description	API for Request Forgot Password
//	@ID				User-ForgotPassword
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			data	body	dto.ForgotPassword	true	"body data"
//	@Success		200
//	@Router			/v1/auth/forgot-password [post]
func (c *AuthControllerHTTP) ForgotPassword(ctx *gin.Context) {

	var input dto.ForgotPassword

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, "Request Invalid "+err.Error(), 400, 0, nil)
		return
	}
	err := c.authUsecase.ForgotPassword(input)

	if err.Error != nil {
		util.APIResponse(ctx, err.Error.Error(), err.Code, 0, nil)
	} else {
		util.APIResponse(ctx, "Request Forgot Password Success", http.StatusOK, 0, nil)
	}
}

// Login
//
//	@Tags			Auth
//	@Summary		Login
//	@Description	API for Login
//	@ID				User-Login
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			data	body	dto.LoginRequest	true	"body data"
//	@Success		200
//	@Router			/v1/auth/login [post]
func (c *AuthControllerHTTP) Login(ctx *gin.Context) {

	var input dto.LoginRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, "Request Invalid "+err.Error(), 400, 0, nil)
		return
	}
	res, err := c.authUsecase.Login(input)

	if err.Error != nil {
		util.APIResponse(ctx, err.Error.Error(), err.Code, 0, nil)
	} else {
		util.APIResponse(ctx, "Login Success", http.StatusOK, 0, res)
	}
}
