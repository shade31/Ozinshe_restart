package core

import (
	"net/http"
	"strconv"

	"Ozinshe_restart/internal/models"

	"github.com/gin-gonic/gin"
)

func (ac *AgeController) UpdateAge(c *gin.Context) {
	ageIDStr := c.Param("ageID")
	ageID, err := strconv.Atoi(ageIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_INVALID_GENRE_ID",
					Message: "Invalid age ID",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
	}

	var age models.Age
	if err := c.ShouldBindJSON(&age); err != nil {
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

	updatedAge, err := ac.AgeRepository.UpdateAge(c, ageID, age)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_UPDATE_AGE",
					Message: "Can't update age",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: updatedAge})
}
