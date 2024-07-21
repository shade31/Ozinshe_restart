package core

import (
	"Ozinshe_restart/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (gc *GenreController) GetGenreByID(c *gin.Context) {
	genreIDStr := c.Param("genreID")
	genreID, err := strconv.Atoi(genreIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_INVALID_GENRE_ID",
					Message: "Invalid genre ID",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
	}
	genres, err := gc.GenreRepository.GetGenreByID(c, genreID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_GENRE",
					Message: "Can't get genre from db",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Result: genres})
}
