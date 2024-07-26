package repository

import (
	"context"
	"errors"
	"time"

	"Ozinshe_restart/internal/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ContentRepository struct {
	db *pgxpool.Pool
}

func NewContentRepository(db *pgxpool.Pool) models.ContentRepository {
	return &ContentRepository{db: db}
}

func (cr *ContentRepository) CreateContent(c context.Context, content models.Content) (int, error) {
	var contentID int
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	var exists bool
	checkQuery := `select exists(select 1 from content where title = $1 AND release_year = $2 AND season = $3 AND episode = $4)`
	err := cr.db.QueryRow(c, checkQuery, content.Title, content.Release_year, content.Season, content.Episode).Scan(&exists)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, errors.New("dublicate entry: a content with same title and release_year already exists")
	}

	if content.Content_type == "Serial" {
		if content.Season == "" || content.Episode == "" {
			return 0, errors.New("season and episode are required for type 'Serial'")
		}
	}
	if content.Content_type == "Movie" {
		if content.Season != "" || content.Episode != "" {
			return 0, errors.New("season and episode must be empty for type 'Movie'")
		}
	}
	userQuery := `INSERT INTO content(
		title, description, release_year, duration, season, episode, content_type, director, producer, genre_id, age_id, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) returning content_id;`
	err = cr.db.QueryRow(c, userQuery, content.Title, content.Description, content.Release_year, content.Duration, content.Season, content.Episode, content.Content_type, content.Director, content.Producer, content.Genre_id, content.Age_id, currentTime).Scan(&contentID)
	if err != nil {
		return 0, err
	}
	return contentID, nil
}

func (cr *ContentRepository) GetContentByID(c context.Context, contentID int) (models.Content, error) {
	content := models.Content{}

	query := `SELECT content_id, title, description, release_year, duration, coalesce(season, ''), coalesce(episode, ''), content_type, director, producer, genre_id, age_id, created_at, coalesce(deleted_at, '') FROM content where content_id = $1`
	row := cr.db.QueryRow(c, query, contentID)
	err := row.Scan(&content.Content_id, &content.Title, &content.Description, &content.Release_year, &content.Duration, &content.Season, &content.Episode, &content.Content_type, &content.Director, &content.Producer, &content.Genre_id, &content.Age_id, &content.Created_at, &content.Deleted_at)

	if err != nil {
		return content, err
	}

	return content, nil
}

func (cr *ContentRepository) GetContentByGenre(c context.Context, genreID int) ([]models.Content, error) {
	var contents []models.Content

	query := `SELECT content_id, title, release_year, content_type, genre_id, age_id, created_at, coalesce(deleted_at, '') FROM content where genre_id = $1`
	rows, err := cr.db.Query(c, query, genreID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var content models.Content
		err := rows.Scan(&content.Content_id, &content.Title, &content.Release_year, &content.Content_type, &content.Genre_id, &content.Age_id, &content.Created_at, &content.Deleted_at)
		if err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contents, nil
}

func (cr *ContentRepository) GetContentByTitle(c context.Context, title string) ([]models.Content, error) {
	var contents []models.Content
	searchTitle := "%" + title + "%"

	query := `SELECT content_id, title, release_year, content_type, genre_id, age_id, created_at, coalesce(deleted_at, '') FROM content where title LIKE $1`
	rows, err := cr.db.Query(c, query, searchTitle)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var content models.Content
		err := rows.Scan(&content.Content_id, &content.Title, &content.Release_year, &content.Content_type, &content.Genre_id, &content.Age_id, &content.Created_at, &content.Deleted_at)
		if err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contents, nil
}

func (gr *ContentRepository) GetAllContents(c context.Context) ([]models.Content, error) {
	var contents []models.Content

	query := `SELECT id, name, created_at, COALESCE(deleted_at, '')  FROM content`
	rows, err := gr.db.Query(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var content models.Content
		err := rows.Scan(&content.Content_id, &content.Title, &content.Created_at, &content.Deleted_at)
		if err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contents, nil
}

func (cr *ContentRepository) UpdateContent(c context.Context, contentID int, content models.Content) (models.Content, error) {
	updatedContent := models.Content{}

	query := `Update content set title = $2, description = $3, release_year = $4, duration = $5, season = $6, episode = $7, content_type = $8, director = $9, producer = $10, genre_id = $11, age_id = $12 where content_id = $1 returning content_id, title, description, release_year, duration, season, episode, content_type, director, producer, genre_id, age_id`

	var (
		content_id   int
		title        string
		description  string
		release_year string
		duration     string
		season       string
		episode      string
		content_type string
		director     string
		producer     string
		genre_id     int
		age_id       int
	)
	err := cr.db.QueryRow(c, query, contentID, content.Title, content.Description, content.Release_year, content.Duration, content.Season, content.Episode, content.Content_type, content.Director, content.Producer, content.Genre_id, content.Age_id).Scan(&content_id, &title, &description, &release_year, &duration, &season, &episode, &content_type, &director, &producer, &genre_id, &age_id)
	if err != nil {
		return updatedContent, err
	}

	updatedContent = models.Content{
		Content_id:   uint(content_id),
		Title:        title,
		Description:  description,
		Release_year: release_year,
		Duration:     duration,
		Season:       season,
		Episode:      episode,
		Content_type: content_type,
		Director:     director,
		Producer:     producer,
		Genre_id:     genre_id,
		Age_id:       age_id,
	}

	return updatedContent, nil
}

func (cr *ContentRepository) DeleteContent(c context.Context, contentID int) (models.Content, error) {
	deletedContent := models.Content{}

	query := `Delete from content where content_id = $1 returning content_id`

	var (
		content_id int
	)
	err := cr.db.QueryRow(c, query, contentID).Scan(&content_id)
	if err != nil {
		return deletedContent, err
	}

	deletedContent = models.Content{
		Content_id: uint(content_id),
	}

	return deletedContent, nil
}
