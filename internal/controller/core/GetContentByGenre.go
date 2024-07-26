package core

import (
	"Ozinshe_restart/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (cc *ContentController) GetContentByGenre(c *gin.Context) {
	genreIDStr := c.Param("genreID")
	genreID, err := strconv.Atoi(genreIDStr)
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
	contents, err := cc.ContentRepository.GetContentByGenre(c, genreID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_GENRE",
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
