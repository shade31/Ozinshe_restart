package core

import (
	"net/http"
	"strconv"

	"Ozinshe_restart/internal/models"

	"github.com/gin-gonic/gin"
)

func (ac *AgeController) DeleteAge(c *gin.Context) {
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

	_, err = ac.AgeRepository.DeleteAge(c, ageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_DELETE_GENRE",
					Message: "Can't delete age",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: "Age deleted succesfully"})
}
