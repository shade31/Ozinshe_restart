package core

import (
	"net/http"
	"strconv"

	"Ozinshe_restart/internal/models"

	"github.com/gin-gonic/gin"
)

func (sc *ScreenshotController) DeleteScreenshot(c *gin.Context) {
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

	_, err = sc.ScreenshotRepository.DeleteScreenshot(c, screenshotID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_DELETE_SCREENSHOT",
					Message: "Can't delete screenshot",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: "Screenshot deleted succesfully"})
}
