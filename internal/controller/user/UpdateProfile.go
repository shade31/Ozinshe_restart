package user

import (
	"net/http"

	"Ozinshe_restart/internal/models"

	"github.com/gin-gonic/gin"
)

func (sc *UserController) UpdateProfile(c *gin.Context) {
	userID := c.GetUint("userID")

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
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

	profile, err := sc.UserRepository.UpdateProfile(c, int(userID), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_UPDATE_USER_PROFILE",
					Message: "Can't update profile",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: profile})
}
