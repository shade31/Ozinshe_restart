package user

import (
	"Ozinshe_restart/internal/controller/auth"
	"Ozinshe_restart/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (sc *UserController) ChangePassword(c *gin.Context) {
	var password models.Password
	userID := c.GetUint("userID")
	if err := c.ShouldBindJSON(&password); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_PARSE_REQUEST_BODY",
					Message: "Can't parse request body",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}

	//Validate password
	if err := auth.ValidatePassword(password.Password); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_INVALID_PASSWORD",
					Message: err.Error(),
				},
			},
		})
		return
	}

	//ChangePassword
	response, err := sc.UserRepository.ChangePassword(c, int(userID), password)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_CHANGE_PASSWORD",
					Message: "Can't change profile password",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Result: response})

}
