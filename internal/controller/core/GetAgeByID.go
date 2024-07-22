package core

import (
	"Ozinshe_restart/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ac *AgeController) GetAgeByID(c *gin.Context) {
	ageIDStr := c.Param("ageID")
	ageID, err := strconv.Atoi(ageIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_INVALID_AGE_ID",
					Message: "Invalid age ID",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
	}
	ages, err := ac.AgeRepository.GetAgeByID(c, ageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_AGE",
					Message: "Can't get age from db",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Result: ages})
}
