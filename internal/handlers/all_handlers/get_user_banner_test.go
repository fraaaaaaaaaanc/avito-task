package allHandlers

import (
	"avito-tech/internal/logger"
	hlModel "avito-tech/internal/models/hanlders_models"
	storageModels "avito-tech/internal/models/storage_model"
	mock "avito-tech/internal/storage/mock"
	"avito-tech/internal/token"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"io"

	gomock "github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestGetUserBanner(t *testing.T) {
	_ = logger.NewZapLogger("", "local")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mock.NewMockStorageBanner(ctrl)
	token := token.NewTokenAccount("testKey")
	hndl, _ := NewHandlers(mockStorage, token)

	gomock.InOrder(
		mockStorage.EXPECT().GetUserBanner(gomock.Any(), hlModel.GetUserBannerModel{
			TagId: 2,
			FeatureId: 1,
			UseLastRevision: false,
		}).Return(nil, storageModels.ErrBannerNotFound),
		mockStorage.EXPECT().GetUserBanner(gomock.Any(), hlModel.GetUserBannerModel{
			TagId: 2,
			FeatureId: 1,
			UseLastRevision: false,
		}).Return(nil, errors.New("error with work database")),
		mockStorage.EXPECT().GetUserBanner(gomock.Any(), hlModel.GetUserBannerModel{
			TagId: 2,
			FeatureId: 1,
			UseLastRevision: false,
		}).Return(&hlModel.BannerContentModel{
			Title: "title",
			Text: "text",
			URL: "url",
		}, nil),
	)

	method := http.MethodGet
	urlLogin := "http://localhost:8080/user_banner"
	type req struct {
		body string
	}
	type resp struct {
		statusCode int
		respBody string
		contentType string
	}
	tests := []struct{
		name string
		req
		resp
	}{
		{
			name: "GET request was sent to \"http://localhost:8080/user_banner\" with an empty request, " +
					"it should return the Status Code 400",
			req: req{
				body: "",
			},
			resp: resp{
				statusCode: http.StatusBadRequest,
				respBody: `{
					"Error": "incorrect data"
				}`,
				contentType: "application/json",
			},
		},
		{
			name: "GET request was sent to \"http://localhost:8080/user_banner\" with invalid parameters in the request (missing the required tag_id parameter), " +
					"it should return the Status Code 400",
			req: req{
				body: "feature_id=1",
			},
			resp: resp{
				statusCode: http.StatusBadRequest,
				respBody: `{
					"Error": "incorrect data"
				}`,
				contentType: "application/json",
			},
		},
		{
			name: "GET request was sent to \"http://localhost:8080/user_banner\" with invalid parameters in the request (inconsistency with the expected data type), " +
					"it should return the Status Code 400",
			req: req{
				body: "feature_id=1&tag_id=2&use_last_revision=7",
			},
			resp: resp{
				statusCode: http.StatusBadRequest,
				respBody: `{
					"Error": "incorrect data"
				}`,
				contentType: "application/json",
			},
		},
		{
			name: "GET request was sent to \"http://localhost:8080/user_banner\" with parameters for which the banner is missing, " +
					"it should return the Status Code 404",
			req: req{
				body: "feature_id=1&tag_id=2&use_last_revision=false",
			},
			resp: resp{
				statusCode: http.StatusNotFound,
			},
		},
		{
			name: "GET request was sent to \"http://localhost:8080/user_banner\" error working with the database, " +
					"it should return the Status Code 500",
			req: req{
				body: "feature_id=1&tag_id=2&use_last_revision=false",
			},
			resp: resp{
				statusCode: http.StatusInternalServerError,
				respBody: `{
					"Error": "internal server error"
				}`,
				contentType: "application/json",
			},
		},
		{
			name: "GET request was sent to \"http://localhost:8080/user_banner\" with the right parameters, " +
					"it should return the Status Code 200",
			req: req{
				body: "feature_id=1&tag_id=2&use_last_revision=false",
			},
			resp: resp{
				statusCode: http.StatusOK,
				respBody: `{
					"title": "title",
					"text": "text",
					"url": "url"
				}`,
				contentType: "application/json",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func (t *testing.T)  {
			request := httptest.NewRequest(method, fmt.Sprintf("%s?%s", urlLogin, test.body), strings.NewReader(test.body))
			w := httptest.NewRecorder()
			hndl.GetUserBannerHandler(w, request)

			resp := w.Result()
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)

			assert.Equal(t, test.statusCode, resp.StatusCode)
			if test.respBody != "" {
				assert.JSONEq(t, test.respBody, string(body))
				assert.Equal(t, test.contentType, resp.Header.Get("Content-Type"))
			}
		})
	}  
}