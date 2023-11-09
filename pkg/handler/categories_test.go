package handler

import (
	"bytes"
	"errors"
	"marketplace"
	"marketplace/pkg/service"
	mock_service "marketplace/pkg/service/mocks"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_getCategoriesAll(t *testing.T) {
	type mockBehavior func(r *mock_service.MockCategories, categories []marketplace.CategoriesList)

	tests := []struct {
		name                 string
		inputBody            string
		inputCategory        []marketplace.CategoriesList
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			inputCategory: []marketplace.CategoriesList{
				{
					Id:    1,
					Title: "Обувь",
				},
				{
					Id:    1,
					Title: "Одежда",
				},
			},
			mockBehavior: func(r *mock_service.MockCategories, category []marketplace.CategoriesList) {
				returnCategories := []marketplace.CategoriesList{
					{
						Id:    1,
						Title: "Обувь",
					},
					{
						Id:    2,
						Title: "Одежда",
					},
				}
				r.EXPECT().GetAllCategories().Return(returnCategories, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"data":[{"id":1,"title":"Обувь"},{"id":2,"title":"Одежда"}]}`,
		},
		{
			name: "Service Error",
			mockBehavior: func(r *mock_service.MockCategories, category []marketplace.CategoriesList) {
				r.EXPECT().GetAllCategories().Return(nil, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockCategories(c)
			test.mockBehavior(repo, test.inputCategory)

			services := &service.Service{Categories: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.GET("/categories", handler.getAllCategories)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/categories",
				bytes.NewBufferString("inputBody"))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_getCategoryById(t *testing.T) {
	type mockBehavior func(r *mock_service.MockCategories, category marketplace.CategoriesList)

	tests := []struct {
		name                 string
		inputBody            string
		uri                  string
		inputCategory        marketplace.CategoriesList
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `"id":1`,
			uri:       "/categories/1",
			inputCategory: marketplace.CategoriesList{
				Id:    1,
				Title: "Обувь",
			},
			mockBehavior: func(r *mock_service.MockCategories, category marketplace.CategoriesList) {
				returnCategory := marketplace.CategoriesList{
					Id:    1,
					Title: "Обувь",
				}
				r.EXPECT().GetById(1).Return(returnCategory, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1,"title":"Обувь"}`,
		},
		{
			name:      "Service Error",
			inputBody: ``,
			uri:       "/categories/1",
			mockBehavior: func(r *mock_service.MockCategories, category marketplace.CategoriesList) {
				returnCategory := marketplace.CategoriesList{
					Id:    1,
					Title: "Adidas",
				}
				r.EXPECT().GetById(1).Return(returnCategory, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
		{
			name:                 "Id input error",
			inputBody:            ``,
			uri:                  "/categories/p",
			mockBehavior:         func(r *mock_service.MockCategories, category marketplace.CategoriesList) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid id param"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockCategories(c)
			test.mockBehavior(repo, test.inputCategory)

			services := &service.Service{Categories: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.GET("/categories/:id", handler.getCategoriesById)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", test.uri,
				bytes.NewBufferString("inputBody"))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
