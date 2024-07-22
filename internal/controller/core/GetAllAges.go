package core

import (
	"Ozinshe_restart/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (gc *AgeController) GetAllAges(c *gin.Context) {

	ages, err := gc.AgeRepository.GetAllAges(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_ALL_AGES",
					Message: "Can't get list of all ages from db",
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
