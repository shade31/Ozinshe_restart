package pkg

import "github.com/jackc/pgx/v4/pgxpool"

type Application struct {
	DB *pgxpool.Pool
}

func App() (Application, error) {
	app := &Application{}
	conn, err := ConnectDatabase()
	if err != nil {
		return Application{}, err
	}
	app.DB = conn

	return *app, nil
}

// func (app *Application) CloseDBConnection() {
// 	app.CloseDBConnection()
// }
