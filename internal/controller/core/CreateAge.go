package core

import (
	"Ozinshe_restart/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AgeController struct {
	AgeRepository models.AgeRepository
}

func (ac *AgeController) CreateAge(c *gin.Context) {
	var request models.Age

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_JSON",
					Message: "Datas dont match with struct of CreateAge",
				},
			},
		})
		return
	}

	_, err := ac.AgeRepository.CreateAge(c, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_CREATE_AGE",
					Message: "Couldn't create age",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Result: "Age created!"})
}
