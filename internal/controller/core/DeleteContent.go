package core

import (
	"net/http"
	"strconv"

	"Ozinshe_restart/internal/models"

	"github.com/gin-gonic/gin"
)

func (cc *ContentController) DeleteContent(c *gin.Context) {
	contentIDStr := c.Param("contentID")
	contentID, err := strconv.Atoi(contentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_INVALID_AGE_ID",
					Message: "Invalid content ID",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
	}

	_, err = cc.ContentRepository.DeleteContent(c, contentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_DELETE_CONTENT",
					Message: "Can't delete content",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: "Content deleted succesfully"})
}
