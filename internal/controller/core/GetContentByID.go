package core

import (
	"Ozinshe_restart/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (cc *ContentController) GetContentByID(c *gin.Context) {
	contentIDStr := c.Param("contentID")
	contentID, err := strconv.Atoi(contentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_INVALID_CONTENT_ID",
					Message: "Invalid content ID",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
	}
	contents, err := cc.ContentRepository.GetContentByID(c, contentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_CONTENT",
					Message: "Can't get content from db",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Result: contents})
}
