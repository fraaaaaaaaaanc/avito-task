// Package app that describes the structure of the application, as well as methods for creating an application object
// and launching it
package app

import (
	"avito-tech/internal/config"
	"net/http"
)

type app struct {
	//TODO описать струткуру
	flagsConf *config.Flags
}

func newApp() (*app, error) {
	//TODO описать создание объекта структуры
	flagsConf := config.ParseConfFlags()

	app := &app{
		flagsConf: flagsConf,
	}

	return app, nil
}

func Run() error {
	//TODO описать метод запуска программы
	_, err := newApp()
	if err != nil {
		return err
	}
	err = http.ListenAndServe()
	return err
}
