package pkg

import (
)

type Application struct {
	DB   DB
}

func App() (Application, error) {
	app := &Application{}
	conn, err := NewConn()
	if err != nil {
		return Application{}, err
	}
	app.DB = conn

	return *app, nil
}

func (app *Application) CloseDBConnection() {
	app.CloseDBConnection()
}
