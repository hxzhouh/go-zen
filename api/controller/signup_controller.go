package controller

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hxzhouh/go-zen.git/bootstrap"
	"github.com/hxzhouh/go-zen.git/domain"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

// Signup 注册
// @Summary 注册用户
// @Description 注册用户，
// @Tags 用户相关接口
// @Param request_id header string true "Request ID"
// @Accept application/json
// @Produce application/json
// @Param object query domain.SignupRequest false "注册查询参数"
// @Success 200 {object} domain.SignupResponse
// @Router /signup [post]
func (sc *SignupController) Signup(c *gin.Context) {
	var request domain.SignupRequest
	err := c.ShouldBind(&request)
	if err != nil {
		slog.Error("SignupController Signup error", "error", err.Error())
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	slog.Info("SignupController", "Signup", "request", request)

	_, err = sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.SignupUsecase.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.SetCookie("accessToken", accessToken, sc.Env.AccessTokenExpiryHour*60*60, "/", "", false, true)
	c.SetCookie("refreshToken", refreshToken, sc.Env.RefreshTokenExpiryHour*60*60, "/", "", false, true)
	c.JSON(http.StatusOK, signupResponse)
}
