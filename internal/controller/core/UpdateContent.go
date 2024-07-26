package core

import (
	"net/http"
	"strconv"

	"Ozinshe_restart/internal/models"

	"github.com/gin-gonic/gin"
)

func (ac *ContentController) UpdateContent(c *gin.Context) {
	contentIDStr := c.Param("contentID")
	contentID, err := strconv.Atoi(contentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_INVALID_GENRE_ID",
					Message: "Invalid content ID",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
	}

	var content models.Content
	if err := c.ShouldBindJSON(&content); err != nil {
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

	updatedContent, err := ac.ContentRepository.UpdateContent(c, contentID, content)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_UPDATE_CONTENT",
					Message: "Can't update content",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: updatedContent})
}
