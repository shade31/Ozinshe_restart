package repository

import (
	"context"
	"time"

	"Ozinshe_restart/internal/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AgeRepository struct {
	db *pgxpool.Pool
}

func NewAgeRepository(db *pgxpool.Pool) models.AgeRepository {
	return &AgeRepository{db: db}
}

func (gr *AgeRepository) CreateAge(c context.Context, age models.Age) (int, error) {
	var ageID int
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	userQuery := `INSERT INTO age(
		name, created_at)
		VALUES ($1, $2) returning id;`
	err := gr.db.QueryRow(c, userQuery, age.Name, currentTime).Scan(&ageID)
	if err != nil {
		return 0, err
	}
	return ageID, nil
}

func (gr *AgeRepository) GetAgeByID(c context.Context, ageID int) (models.Age, error) {
	age := models.Age{}

	query := `SELECT id, name, created_at, coalesce(deleted_at, '') FROM age where id = $1`
	row := gr.db.QueryRow(c, query, ageID)
	err := row.Scan(&age.Id, &age.Name, &age.Created_at, &age.Deleted_at)

	if err != nil {
		return age, err
	}

	return age, nil
}

func (gr *AgeRepository) GetAllAges(c context.Context) ([]models.Age, error) {
	var ages []models.Age

	query := `SELECT id, name, created_at, COALESCE(deleted_at, '')  FROM age`
	rows, err := gr.db.Query(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var age models.Age
		err := rows.Scan(&age.Id, &age.Name, &age.Created_at, &age.Deleted_at)
		if err != nil {
			return nil, err
		}
		ages = append(ages, age)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ages, nil
}

func (gr *AgeRepository) UpdateAge(c context.Context, ageID int, age models.Age) (models.Age, error) {
	updatedAge := models.Age{}

	query := `Update age set name = $2 where id = $1 returning id, name`

	var (
		id   int
		name string
	)
	err := gr.db.QueryRow(c, query, ageID, age.Name).Scan(&id, &name)
	if err != nil {
		return updatedAge, err
	}

	updatedAge = models.Age{
		Id:   uint(id),
		Name: name,
	}

	return updatedAge, nil
}

func (gr *AgeRepository) DeleteAge(c context.Context, ageID int) (models.Age, error) {
	deletedAge := models.Age{}

	query := `Delete from age where id = $1 returning id`

	var (
		id int
	)
	err := gr.db.QueryRow(c, query, ageID).Scan(&id)
	if err != nil {
		return deletedAge, err
	}

	deletedAge = models.Age{
		Id: uint(id),
	}

	return deletedAge, nil
}
