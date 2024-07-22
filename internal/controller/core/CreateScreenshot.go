package core

import (
	"Ozinshe_restart/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ScreenshotController struct {
	ScreenshotRepository models.ScreenshotRepository
}

func (ac *ScreenshotController) CreateScreenshot(c *gin.Context) {
	var request models.Screenshot

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_JSON",
					Message: "Datas dont match with struct of CreateScreenshot",
				},
			},
		})
		return
	}

	screenshotID, err := ac.ScreenshotRepository.CreateScreenshot(c, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_CREATE_GENRE",
					Message: "Couldn't create screenshot",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Result: screenshotID})
}
