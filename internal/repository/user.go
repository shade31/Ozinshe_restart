package repository

import (
	"context"
	"fmt"
	"time"

	"Ozinshe_restart/internal/models"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) models.UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(c context.Context, user models.UserRequest) (int, error) {
	var userID int
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	userQuery := `INSERT INTO users(
		email, password, created_at)
		VALUES ($1, $2, $3) returning id;`
	err := ur.db.QueryRow(c, userQuery, user.Email, user.Password, currentTime).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (ur *UserRepository) GetUserByEmail(c context.Context, email string) (models.User, error) {
	user := models.User{}
	fmt.Println(email)
	query := `SELECT id, email, password, coalesce(name , ''), coalesce(birthdate , ''), coalesce(phone , ''), created_at FROM users where email = $1`
	row := ur.db.QueryRow(c, query, email)
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Birthdate, &user.Phone, &user.Created_at)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByID(c context.Context, userID int) (models.User, error) {
	user := models.User{}

	query := `SELECT id, email, password, coalesce(name , ''), coalesce(birthdate , ''), coalesce(phone , ''), created_at FROM users where id = $1`
	row := ur.db.QueryRow(c, query, userID)
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Birthdate, &user.Phone, &user.Created_at)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) GetProfile(c context.Context, userID int) (models.User, error) {
	user := models.User{}

	query := `SELECT id, email, password, coalesce(name , ''), coalesce(birthdate , ''), coalesce(phone , ''), created_at FROM users where id = $1`
	row := ur.db.QueryRow(c, query, userID)
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Birthdate, &user.Phone, &user.Created_at)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) UpdateProfile(c context.Context, userID int, user models.User) (models.User, error) {
	updatedUser := models.User{}

	query := `Update users set email = $2, name = $3, phone = $4, birthdate = $5 where id = $1 returning id, email, name, phone, birthdate`

	var (
		id        int
		email     string
		name      string
		phone     string
		birthdate string
	)
	err := ur.db.QueryRow(c, query, userID, user.Email, user.Name, user.Phone, user.Birthdate).Scan(&id, &email, &name, &phone, &birthdate)
	if err != nil {
		return updatedUser, err
	}

	updatedUser = models.User{
		Id:        uint(id),
		Email:     email,
		Name:      name,
		Phone:     phone,
		Birthdate: birthdate,
	}

	return updatedUser, nil
}

func (ur *UserRepository) ChangePassword(c context.Context, userID int, p models.Password) (models.UserIDResponse, error) {
	var response models.UserIDResponse
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return response, fmt.Errorf("error hashing password: %w", err)
	}

	query := `Update users set password = $2 where id = $1 returning id`
	err = ur.db.QueryRow(c, query, userID, string(hashedPassword)).Scan(&response.Id)
	if err != nil {
		return models.UserIDResponse{}, err
	}
	return response, nil
}
