package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/username/GitRepoName/internal/controller/tokenutil"
	"github.com/username/GitRepoName/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (lc *AuthController) Signin(c *gin.Context) {
	var loginRequest models.LoginRequest

	err := c.ShouldBind(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_JSON",
					Message: "Datas dont match with struct of signin",
				},
			},
		})
		return
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "EMPTY_VALUES",
					Message: "Empty values are written in the form",
				},
			},
		})
		return
	}

	user, err := lc.UserRepository.GetUserByEmail(c, loginRequest.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_USER",
					Message: "User with this email doesn't found",
				},
			},
		})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)) != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "PASSWORD_INCORRECT",
					Message: "Password doesn't match",
				},
			},
		})
		return
	}
	accessToken, err := tokenutil.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "TOKEN_ERROR",
					Message: "Error to create access token",
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: accessToken})
}
