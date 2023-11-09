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

func TestHandler_getProductAll(t *testing.T) {
	type mockBehavior func(r *mock_service.MockProducts, products []marketplace.ProductList)

	tests := []struct {
		name                 string
		inputBody            string
		inputProdycts        []marketplace.ProductList
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			inputProdycts: []marketplace.ProductList{
				{
					Id:          1,
					Title:       "Лофферы синие",
					Description: "",
					Price:       "10.2",
					Brand:       "1",
					Category:    "1",
				},
				{
					Id:          2,
					Title:       "Лофферы красные",
					Description: "",
					Price:       "10.2",
					Brand:       "1",
					Category:    "1",
				},
				{
					Id:          3,
					Title:       "Куртка синяя",
					Description: "",
					Price:       "12.2",
					Brand:       "1",
					Category:    "1",
				},
			},
			mockBehavior: func(r *mock_service.MockProducts, products []marketplace.ProductList) {
				returnCategories := []marketplace.ProductList{
					{
						Id:          1,
						Title:       "Лофферы синие",
						Description: "",
						Price:       "10.2",
						Brand:       "1",
						Category:    "1",
					},
					{
						Id:          2,
						Title:       "Лофферы красные",
						Description: "",
						Price:       "10.2",
						Brand:       "1",
						Category:    "1",
					},
					{
						Id:          3,
						Title:       "Куртка синяя",
						Description: "",
						Price:       "12.2",
						Brand:       "1",
						Category:    "1",
					},
				}
				r.EXPECT().GetAll().Return(returnCategories, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"data":[{"id":1,"title":"Лофферы синие","description":"","price":"10.2","brand":"1","category":"1"},{"id":2,"title":"Лофферы красные","description":"","price":"10.2","brand":"1","category":"1"},{"id":3,"title":"Куртка синяя","description":"","price":"12.2","brand":"1","category":"1"}]}`,
		},
		{
			name: "Service Error",
			mockBehavior: func(r *mock_service.MockProducts, products []marketplace.ProductList) {
				r.EXPECT().GetAll().Return(nil, errors.New("something went wrong"))
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

			repo := mock_service.NewMockProducts(c)
			test.mockBehavior(repo, test.inputProdycts)

			services := &service.Service{Products: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.GET("/products", handler.getProducts)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/products",
				bytes.NewBufferString("inputBody"))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_getProductById(t *testing.T) {
	type mockBehavior func(r *mock_service.MockProducts, product marketplace.ProductList)

	tests := []struct {
		name                 string
		inputBody            string
		uri                  string
		inputProduct         marketplace.ProductList
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `"id":1`,
			uri:       "/products/1",
			inputProduct: marketplace.ProductList{
				Id:          1,
				Title:       "Куртка синяя",
				Description: "",
				Price:       "12.2",
				Brand:       "1",
				Category:    "1",
			},
			mockBehavior: func(r *mock_service.MockProducts, product marketplace.ProductList) {
				returnProduct := marketplace.ProductList{
					Id:          3,
					Title:       "Куртка синяя",
					Description: "",
					Price:       "12.2",
					Brand:       "1",
					Category:    "1",
				}
				r.EXPECT().GetById(1).Return(returnProduct, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":3,"title":"Куртка синяя","description":"","price":"12.2","brand":"1","category":"1"}`,
		},
		{
			name:      "Service Error",
			inputBody: ``,
			uri:       "/products/1",
			mockBehavior: func(r *mock_service.MockProducts, category marketplace.ProductList) {
				returnProduct := marketplace.ProductList{
					Id:          3,
					Title:       "Куртка синяя",
					Description: "",
					Price:       "12.2",
					Brand:       "1",
					Category:    "1",
				}
				r.EXPECT().GetById(1).Return(returnProduct, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
		{
			name:                 "Id input error",
			inputBody:            ``,
			uri:                  "/products/p",
			mockBehavior:         func(r *mock_service.MockProducts, product marketplace.ProductList) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid id param"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockProducts(c)
			test.mockBehavior(repo, test.inputProduct)

			services := &service.Service{Products: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.GET("/products/:id", handler.getProductById)

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
