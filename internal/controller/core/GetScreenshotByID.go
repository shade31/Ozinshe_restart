package core

import (
	"Ozinshe_restart/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (sc *ScreenshotController) GetScreenshotByID(c *gin.Context) {
	screenshotIDStr := c.Param("screenshotID")
	screenshotID, err := strconv.Atoi(screenshotIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_INVALID_SCREENSHOT_ID",
					Message: "Invalid screenshot ID",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
	}
	screenshots, err := sc.ScreenshotRepository.GetScreenshotByID(c, screenshotID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_SCREENSHOT",
					Message: "Can't get screenshot from db",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Result: screenshots})
}
