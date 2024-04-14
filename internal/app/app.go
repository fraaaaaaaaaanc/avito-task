// Package app that describes the structure of the application, as well as methods for creating an application object
// and launching it
package app

import (
	"avito-tech/internal/caches"
	"avito-tech/internal/config"
	"avito-tech/internal/handlers/all_handlers"
	"avito-tech/internal/logger"
	rt "avito-tech/internal/router"
	"avito-tech/internal/storage"
	postgreStorage "avito-tech/internal/storage/postgre_storage"
	"avito-tech/internal/token"
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type app struct {
	//TODO описать струткуру
	flagsConf    *config.Flags
	hl           *allHandlers.Handlers
	strg         storage.StorageBanner
	tokenAccount *token.TokenAccount
	router       chi.Router
	redis *caches.Cache
	// delQueryChan chan storageModels.DeleteBannerFeatureOrTagModel
}

func newApp() (*app, error) {
	flagsConf := config.ParseConfFlags()

	err := logger.NewZapLogger(flagsConf.LogFilePath, flagsConf.ProjLvl)
	if err != nil {
		return nil, err
	}

	strg, err := postgreStorage.NewStorage(flagsConf.DataBaseURI)
	if err != nil {
		logger.Error("error creating the storage object", zap.Error(err))
		return nil, err
	}

	tokenAccount := token.NewTokenAccount(flagsConf.SecretKeyJWTToken)

	hl, err := allHandlers.NewHandlers(strg, tokenAccount)
	if err != nil {
		logger.Error("Error creating the hl object", zap.Error(err))
		return nil, err
	}

	redis, err := caches.NewRedis(flagsConf.RedisURL)
	if err != nil {
		logger.Error("Error creating the redis object", zap.Error(err))
		return nil, err
	}

	router, err := rt.NewRouter(hl, tokenAccount, redis)
	if err != nil {
		logger.Error("error creating the Router object", zap.Error(err))
		return nil, err
	}

	app := &app{
		flagsConf:    flagsConf,
		strg:         strg,
		hl:           hl,
		tokenAccount: tokenAccount,
		router:       router,
		redis: redis,
	}

	return app, nil
}

func AppRun() error {
	app, err := newApp()

	defer logger.CloseFile()
	defer app.strg.CloseDB()

	if err != nil {
		return err
	}
	logger.Info("Server start", zap.String("Running server on", app.flagsConf.String()))
	err = http.ListenAndServe(app.flagsConf.String(), app.router)
	logger.Error("Error run server", zap.Error(err))
	return err
}
