package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hxzhouh/go-zen.git/bootstrap"
	"github.com/hxzhouh/go-zen.git/domain"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

// Login 登录
// @Summary 用户登录接口
// @Description 注册用户，
// @Tags 用户相关接口
// @Param request_id header string true "Request ID"
// @Accept application/json
// @Produce application/json
// @Param object body domain.LoginRequest false "用户登录"
// @Success 200 {object} domain.LoginResponse
// @Router /login [post]
func (lc *LoginController) Login(c *gin.Context) {
	var request domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	user := domain.User{}
	if !isEmail(request.User) {
		user, err = lc.LoginUsecase.GetUserByEmail(c, request.User)
		if err != nil {
			c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given email"})
			return
		}
	} else {
		user, err = lc.LoginUsecase.GetUserByEmail(c, request.User)
		if err != nil {
			c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given username"})
			return
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}

func isEmail(email string) bool {
	// Define a regular expression for validating an email address.
	const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
