package repository

import (
	"context"
	"time"

	"Ozinshe_restart/internal/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ScreenshotRepository struct {
	db *pgxpool.Pool
}

func NewScreenshotRepository(db *pgxpool.Pool) models.ScreenshotRepository {
	return &ScreenshotRepository{db: db}
}

func (sr *ScreenshotRepository) CreateScreenshot(c context.Context, screenshot models.Screenshot) (int, error) {
	var screenshotID int
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	userQuery := `INSERT INTO screenshot(
		content_id, screen, created_at)
		VALUES ($1, $2, $3) returning id;`
	err := sr.db.QueryRow(c, userQuery, screenshot.Content_id, screenshot.Screen, currentTime).Scan(&screenshotID)
	if err != nil {
		return 0, err
	}
	return screenshotID, nil
}

func (sr *ScreenshotRepository) GetScreenshotByID(c context.Context, screenshotID int) ([]models.Screenshot, error) {
	var screenshots []models.Screenshot

	query := `SELECT id, content_id, screen, created_at, coalesce(deleted_at, '') FROM screenshot where content_id = $1`
	rows, err := sr.db.Query(c, query, screenshotID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var screenshot models.Screenshot
		err := rows.Scan(&screenshot.Id, &screenshot.Content_id, &screenshot.Screen, &screenshot.Created_at, &screenshot.Deleted_at)
		if err != nil {
			return nil, err
		}
		screenshots = append(screenshots, screenshot)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return screenshots, nil
}

func (sr *ScreenshotRepository) DeleteScreenshot(c context.Context, screenshotID int) (models.Screenshot, error) {
	deletedScreenshot := models.Screenshot{}

	query := `Delete from screenshot where id = $1 returning id`

	var (
		id int
	)
	err := sr.db.QueryRow(c, query, screenshotID).Scan(&id)
	if err != nil {
		return deletedScreenshot, err
	}

	deletedScreenshot = models.Screenshot{
		Id: uint(id),
	}

	return deletedScreenshot, nil
}
