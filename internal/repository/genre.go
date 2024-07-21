package repository

import (
	"context"
	"time"

	"Ozinshe_restart/internal/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type GenreRepository struct {
	db *pgxpool.Pool
}

func NewGenreRepository(db *pgxpool.Pool) models.GenreRepository {
	return &GenreRepository{db: db}
}

func (gr *GenreRepository) CreateGenre(c context.Context, genre models.Genre) (int, error) {
	var genreID int
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	userQuery := `INSERT INTO genre(
		name, created_at)
		VALUES ($1, $2) returning id;`
	err := gr.db.QueryRow(c, userQuery, genre.Name, currentTime).Scan(&genreID)
	if err != nil {
		return 0, err
	}
	return genreID, nil
}

func (gr *GenreRepository) GetGenreByID(c context.Context, genreID int) (models.Genre, error) {
	genre := models.Genre{}

	query := `SELECT id, name, created_at, coalesce(deleted_at, '') FROM genre where id = $1`
	row := gr.db.QueryRow(c, query, genreID)
	err := row.Scan(&genre.Id, &genre.Name, &genre.Created_at, &genre.Deleted_at)

	if err != nil {
		return genre, err
	}

	return genre, nil
}

func (gr *GenreRepository) GetAllGenres(c context.Context) ([]models.Genre, error) {
	var genres []models.Genre

	query := `SELECT id, name, created_at, COALESCE(deleted_at, '')  FROM genre`
	rows, err := gr.db.Query(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var genre models.Genre
		err := rows.Scan(&genre.Id, &genre.Name, &genre.Created_at, &genre.Deleted_at)
		if err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return genres, nil
}

func (gr *GenreRepository) UpdateGenre(c context.Context, genreID int, genre models.Genre) (models.Genre, error) {
	updatedGenre := models.Genre{}

	query := `Update genre set name = $2 where id = $1 returning id, name`

	var (
		id   int
		name string
	)
	err := gr.db.QueryRow(c, query, genreID, genre.Name).Scan(&id, &name)
	if err != nil {
		return updatedGenre, err
	}

	updatedGenre = models.Genre{
		Id:   uint(id),
		Name: name,
	}

	return updatedGenre, nil
}

func (gr *GenreRepository) DeleteGenre(c context.Context, genreID int) (models.Genre, error) {
	deletedGenre := models.Genre{}

	query := `Delete from genre where id = $1 returning id`

	var (
		id int
	)
	err := gr.db.QueryRow(c, query, genreID).Scan(&id)
	if err != nil {
		return deletedGenre, err
	}

	deletedGenre = models.Genre{
		Id: uint(id),
	}

	return deletedGenre, nil
}
