package models

import (
	"context"
)

type User struct {
	Id         uint   `json:"-" db:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Birthdate  string `json:"birthdate"`
	Phone      string `json:"phone"`
	Created_at string `json:"created_at"`
	Deleted_at string `json:"deleted_at"`
	Is_admin   bool   `json:"is_admin"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserIDResponse struct {
	Id uint `json:"id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Password struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserRepository interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	GetUserByID(c context.Context, userID int) (User, error)
	GetProfile(c context.Context, userID int) (User, error)
	UpdateProfile(c context.Context, userID int, u User) (User, error)
	ChangePassword(c context.Context, userID int, p Password) (UserIDResponse, error)
	CreateUser(c context.Context, user UserRequest) (int, error)
}

type Genre struct {
	Id         uint   `json:"id" db:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Deleted_at string `json:"deleted_at"`
}

type GenreRepository interface {
	GetAllGenres(c context.Context) ([]Genre, error)
	GetGenreByID(c context.Context, genreID int) (Genre, error)
	CreateGenre(c context.Context, genre Genre) (int, error)
	UpdateGenre(c context.Context, genreID int, g Genre) (Genre, error)
	DeleteGenre(c context.Context, genreID int) (Genre, error)
}

type Age struct {
	Id         uint   `json:"id" db:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Deleted_at string `json:"deleted_at"`
}

type AgeRepository interface {
	GetAllAges(c context.Context) ([]Age, error)
	GetAgeByID(c context.Context, ageID int) (Age, error)
	CreateAge(c context.Context, age Age) (int, error)
	UpdateAge(c context.Context, ageID int, g Age) (Age, error)
	DeleteAge(c context.Context, ageID int) (Age, error)
}

type Screenshot struct {
	Id         uint   `json:"id" db:"id"`
	Content_id int    `json:"content_id"`
	Screen     string `json:"screen"`
	Created_at string `json:"created_at"`
	Deleted_at string `json:"deleted_at"`
}

type ScreenshotRepository interface {
	GetScreenshotByID(c context.Context, screenshotID int) ([]Screenshot, error)
	CreateScreenshot(c context.Context, screenshot Screenshot) (int, error)
	DeleteScreenshot(c context.Context, screenshotID int) (Screenshot, error)
}
