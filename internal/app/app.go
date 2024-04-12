// Package app that describes the structure of the application, as well as methods for creating an application object
// and launching it
package app

import (
	"avito-tech/internal/config"
	"avito-tech/internal/handlers/all_handlers"
	"avito-tech/internal/logger"
	storageModels "avito-tech/internal/models/storage_model"
	rt "avito-tech/internal/router"
	"avito-tech/internal/storage"
	postgreStorage "avito-tech/internal/storage/postgre_storage"
	"avito-tech/internal/token"
	"context"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type app struct {
	//TODO описать струткуру
	flagsConf    *config.Flags
	hl           *allHandlers.Handlers
	strg         storage.StorageBanner
	tokenAccount *token.TokenAccount
	router       chi.Router
	delQueryChan chan storageModels.DelFeatureOrTagChan
}

func newApp() (*app, error) {
	//TODO описать создание объекта структуры
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

	router, err := rt.NewRouter(hl, tokenAccount)
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
	}

	return app, nil
}

func (a *app) FlushDelete() {
	ticker := time.NewTicker(10 * time.Second)

	var requests []storageModels.DelFeatureOrTagChan

	for {
		select {
		case request := <-a.delQueryChan:
			requests = append(requests, request)
		case <-ticker.C:
			if len(requests) == 0 {
				continue
			}
			go func(requests []storageModels.DelFeatureOrTagChan) {
				var delBanners []storageModels.DelFeatureOrTagChan

				for _, objBanner := range requests {
					delBanners = append(delBanners, objBanner)
				}

				err := a.strg.DelBannerFeatureOrTag(context.Background(), delBanners)
				if err != nil {
					//logger.Error("cannot delete:%v", err)
				}
			}(requests)
			requests = nil
		}
	}
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
