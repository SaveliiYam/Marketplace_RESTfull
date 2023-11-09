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

func TestHandler_getBrandsAll(t *testing.T) {
	type mockBehavior func(r *mock_service.MockBrands, brands []marketplace.BrandsList)

	tests := []struct {
		name                 string
		inputBody            string
		inputBrand           []marketplace.BrandsList
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			inputBrand: []marketplace.BrandsList{
				{
					Id:          1,
					Title:       "Adidas",
					Description: "brand",
				},
				{
					Id:          1,
					Title:       "Nike",
					Description: "brand",
				},
			},
			mockBehavior: func(r *mock_service.MockBrands, brand []marketplace.BrandsList) {
				returnBrands := []marketplace.BrandsList{
					{
						Id:          1,
						Title:       "Adidas",
						Description: "brand",
					},
					{
						Id:          2,
						Title:       "Nike",
						Description: "brand",
					},
				}
				r.EXPECT().GetAllBrands().Return(returnBrands, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"data":[{"id":1,"title":"Adidas","description":"brand"},{"id":2,"title":"Nike","description":"brand"}]}`,
		},
		{
			name: "Service Error",
			mockBehavior: func(r *mock_service.MockBrands, brands []marketplace.BrandsList) {
				r.EXPECT().GetAllBrands().Return(nil, errors.New("something went wrong"))
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

			repo := mock_service.NewMockBrands(c)
			test.mockBehavior(repo, test.inputBrand)

			services := &service.Service{Brands: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.GET("/brands", handler.getAllBrands)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/brands",
				bytes.NewBufferString("inputBody"))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_getBrandById(t *testing.T) {
	type mockBehavior func(r *mock_service.MockBrands, brand marketplace.BrandsList)

	tests := []struct {
		name                 string
		inputBody            string
		uri                  string
		inputBrand           marketplace.BrandsList
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `"id":1`,
			uri:       "/brands/1",
			inputBrand: marketplace.BrandsList{
				Id:          1,
				Title:       "Adidas",
				Description: "brand",
			},
			mockBehavior: func(r *mock_service.MockBrands, brand marketplace.BrandsList) {
				returnBrand := marketplace.BrandsList{
					Id:          1,
					Title:       "Adidas",
					Description: "brand",
				}
				r.EXPECT().GetById(1).Return(returnBrand, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1,"title":"Adidas","description":"brand"}`,
		},
		{
			name:      "Service Error",
			inputBody: ``,
			uri:       "/brands/1",
			mockBehavior: func(r *mock_service.MockBrands, brand marketplace.BrandsList) {
				returnBrand := marketplace.BrandsList{
					Id:          1,
					Title:       "Adidas",
					Description: "brand",
				}
				r.EXPECT().GetById(1).Return(returnBrand, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
		{
			name:                 "Id input error",
			inputBody:            ``,
			uri:                  "/brands/p",
			mockBehavior:         func(r *mock_service.MockBrands, brand marketplace.BrandsList) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid id param"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockBrands(c)
			test.mockBehavior(repo, test.inputBrand)

			services := &service.Service{Brands: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.GET("/brands/:id", handler.getById)

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
