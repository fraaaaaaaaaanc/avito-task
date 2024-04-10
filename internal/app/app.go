// Package app that describes the structure of the application, as well as methods for creating an application object
// and launching it
package app

import (
	"avito-tech/internal/config"
	"avito-tech/internal/handlers/all_handlers"
	"avito-tech/internal/logger"
	rt "avito-tech/internal/router"
	mainStorage "avito-tech/internal/storage/main_storage"
	"avito-tech/internal/token"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
)

type app struct {
	//TODO описать струткуру
	flagsConf    *config.Flags
	hl           *allHandlers.Handlers
	strg         *mainStorage.Storage
	tokenAccount *token.TokenAccount
	router       chi.Router
}

func newApp() (*app, error) {
	//TODO описать создание объекта структуры
	flagsConf := config.ParseConfFlags()

	err := logger.NewZapLogger(flagsConf.LogFilePath, flagsConf.ProjLvl)
	if err != nil {
		return nil, err
	}

	strg, err := mainStorage.NewStorage(flagsConf.DataBaseURI)
	if err != nil {
		logger.Error("error creating the storage object", zap.Error(err))
		return nil, err
	}

	tokenAccount := token.NewTokenAccount(flagsConf.SecretKeyJWTToken)

	hl := allHandlers.NewHandlers(strg, tokenAccount)

	router, err := rt.NewRouter(hl, tokenAccount)
	if err != nil {
		logger.Error("error creating the Router object", zap.Error(err))
		return nil, err
	}

	app := &app{
		flagsConf: flagsConf,
		//strg:      strg,
		hl:           hl,
		tokenAccount: tokenAccount,
		router:       router,
	}

	return app, nil
}

func AppRun() error {
	//TODO описать метод запуска программы
	app, err := newApp()

	//defer logger.CloseFile()
	//defer app.strgs.CloseDB()

	if err != nil {
		return err
	}
	logger.Info("Server start", zap.String("Running server on", app.flagsConf.String()))
	err = http.ListenAndServe(app.flagsConf.String(), app.router)
	logger.Error("Error run server", zap.Error(err))
	return err
}
