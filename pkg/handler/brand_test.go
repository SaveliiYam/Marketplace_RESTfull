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

func TestHanler_createBrand(t *testing.T) {
	type mockBehavior func(r *mock_service.MockBrands, brand marketplace.BrandsList)
	type authBehavior func(r *mock_service.MockAuthorization, user marketplace.User)
	type tokenBehavior func(r *mock_service.MockAuthorization, token string)

	tests := []struct {
		name string

		inputAuthBody string
		inputUser     marketplace.User
		authBehavior  authBehavior

		headerName    string
		headerValue   string
		token         string
		id            int
		tokenBehavior tokenBehavior

		inputBrandBody string
		inputBrand     marketplace.BrandsList
		mockBehavior   mockBehavior

		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:          "OK",
			inputAuthBody: `{"username": "username", "name": "name", "password": "password", "status":true}`,
			inputUser:     marketplace.User{Username: "username", Name: "name", Password: "password", Status: true},
			authBehavior: func(r *mock_service.MockAuthorization, user marketplace.User) {
				r.EXPECT().GenerateToken(user.Username, user.Password).Return("64666873727468626378767a6668735baa61e4c9b93f3f0682250b6cf8331b7ee68fd8", nil)
				r.EXPECT().CreateUser(user).Return(1, nil)
			},

			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			id:          1,
			tokenBehavior: func(r *mock_service.MockAuthorization, token string) {

				r.EXPECT().ParseToken(token).Return(1, nil)
				r.EXPECT().CheckStatus(1).Return(true, nil)
			},

			inputBrandBody: `{"title":"Adidas", "description":"brand"}`,
			inputBrand:     marketplace.BrandsList{Title: "Adidas", Description: "brand"},
			mockBehavior: func(r *mock_service.MockBrands, brand marketplace.BrandsList) {
				r.EXPECT().GetByString(brand.Title).Return(1, errors.New(""))
				r.EXPECT().Create(brand).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `%!s(bool=true){"id":1}`,
		},
		{
			name:          "Invalid brand",
			inputAuthBody: `{"username": "username", "name": "name", "password": "password", "status":true}`,
			inputUser:     marketplace.User{Username: "username", Name: "name", Password: "password", Status: true},
			authBehavior: func(r *mock_service.MockAuthorization, user marketplace.User) {
				r.EXPECT().GenerateToken(user.Username, user.Password).Return("64666873727468626378767a6668735baa61e4c9b93f3f0682250b6cf8331b7ee68fd8", nil)
				r.EXPECT().CreateUser(user).Return(1, nil)
			},

			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			id:          1,
			tokenBehavior: func(r *mock_service.MockAuthorization, token string) {

				r.EXPECT().ParseToken(token).Return(1, nil)
				r.EXPECT().CheckStatus(1).Return(true, nil)
			},

			inputBrandBody: `{"description":"brand"}`,
			inputBrand:     marketplace.BrandsList{Description: "brand"},
			mockBehavior: func(r *mock_service.MockBrands, brand marketplace.BrandsList) {
			},
			expectedStatusCode:   200,
			expectedResponseBody: `%!s(bool=true){"message":"invalid brand param"}`,
		},
		{
			name:          "Already exist",
			inputAuthBody: `{"username": "username", "name": "name", "password": "password", "status":true}`,
			inputUser:     marketplace.User{Username: "username", Name: "name", Password: "password", Status: true},
			authBehavior: func(r *mock_service.MockAuthorization, user marketplace.User) {
				r.EXPECT().GenerateToken(user.Username, user.Password).Return("64666873727468626378767a6668735baa61e4c9b93f3f0682250b6cf8331b7ee68fd8", nil)
				r.EXPECT().CreateUser(user).Return(1, nil)
			},

			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			id:          1,
			tokenBehavior: func(r *mock_service.MockAuthorization, token string) {

				r.EXPECT().ParseToken(token).Return(1, nil)
				r.EXPECT().CheckStatus(1).Return(true, nil)
			},

			inputBrandBody: `{"title":"Adidas", "description":"brand"}`,
			inputBrand:     marketplace.BrandsList{Title: "Adidas", Description: "brand"},
			mockBehavior: func(r *mock_service.MockBrands, brand marketplace.BrandsList) {
				r.EXPECT().GetByString(brand.Title).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `%!s(bool=true){"message":"already exist"}`,
		},
		{
			name:          "Create service error",
			inputAuthBody: `{"username": "username", "name": "name", "password": "password", "status":true}`,
			inputUser:     marketplace.User{Username: "username", Name: "name", Password: "password", Status: true},
			authBehavior: func(r *mock_service.MockAuthorization, user marketplace.User) {
				r.EXPECT().GenerateToken(user.Username, user.Password).Return("64666873727468626378767a6668735baa61e4c9b93f3f0682250b6cf8331b7ee68fd8", nil)
				r.EXPECT().CreateUser(user).Return(1, nil)
			},

			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			id:          1,
			tokenBehavior: func(r *mock_service.MockAuthorization, token string) {

				r.EXPECT().ParseToken(token).Return(1, nil)
				r.EXPECT().CheckStatus(1).Return(true, nil)
			},

			inputBrandBody: `{"title":"Adidas", "description":"brand"}`,
			inputBrand:     marketplace.BrandsList{Title: "Adidas", Description: "brand"},
			mockBehavior: func(r *mock_service.MockBrands, brand marketplace.BrandsList) {
				r.EXPECT().GetByString(brand.Title).Return(1, errors.New(""))
				r.EXPECT().Create(brand).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   200,
			expectedResponseBody: `%!s(bool=true){"message":"something went wrong"}`,
		},
		{
			name:          "Have Not rights",
			inputAuthBody: `{"username": "username", "name": "name", "password": "password"}`,
			inputUser:     marketplace.User{Username: "username", Name: "name", Password: "password"},
			authBehavior: func(r *mock_service.MockAuthorization, user marketplace.User) {
				r.EXPECT().GenerateToken(user.Username, user.Password).Return("64666873727468626378767a6668735baa61e4c9b93f3f0682250b6cf8331b7ee68fd8", nil)
				r.EXPECT().CreateUser(user).Return(1, nil)
				//r.EXPECT().CheckStatus(1).Return(false, errors.New(""))
			},

			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			id:          1,
			tokenBehavior: func(r *mock_service.MockAuthorization, token string) {

				r.EXPECT().ParseToken(token).Return(1, nil)
				r.EXPECT().CheckStatus(1).Return(false, nil)
			},

			inputBrandBody: `{"title":"Adidas", "description":"brand"}`,
			inputBrand:     marketplace.BrandsList{Title: "Adidas", Description: "brand"},
			mockBehavior: func(r *mock_service.MockBrands, brand marketplace.BrandsList) {

			},
			expectedStatusCode:   200,
			expectedResponseBody: `%!s(bool=false){"message":"you do not have sufficient rights"}`,
		},
		{
			name:          "Invalid input",
			inputAuthBody: `{"username": "username", "name": "name", "password": "password"}`,
			inputUser:     marketplace.User{Username: "username", Name: "name", Password: "password"},
			authBehavior: func(r *mock_service.MockAuthorization, user marketplace.User) {
				r.EXPECT().GenerateToken(user.Username, user.Password).Return("64666873727468626378767a6668735baa61e4c9b93f3f0682250b6cf8331b7ee68fd8", nil)
				r.EXPECT().CreateUser(user).Return(1, nil)
				//r.EXPECT().CheckStatus(1).Return(false, errors.New(""))
			},

			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			id:          1,
			tokenBehavior: func(r *mock_service.MockAuthorization, token string) {

				r.EXPECT().ParseToken(token).Return(1, nil)
				r.EXPECT().CheckStatus(1).Return(false, errors.New("something went wrong"))
			},

			inputBrandBody: `{"title":"Adidas", "description":"brand"}`,
			inputBrand:     marketplace.BrandsList{Title: "Adidas", Description: "brand"},
			mockBehavior: func(r *mock_service.MockBrands, brand marketplace.BrandsList) {

			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
		{
			name:          "Invalid authorization",
			inputAuthBody: `{"username": "username", "name": "name", "password": "password"}`,
			inputUser:     marketplace.User{Username: "username", Name: "name", Password: "password"},
			authBehavior: func(r *mock_service.MockAuthorization, user marketplace.User) {
				r.EXPECT().GenerateToken(user.Username, user.Password).Return("64666873727468626378767a6668735baa61e4c9b93f3f0682250b6cf8331b7ee68fd8", nil)
				r.EXPECT().CreateUser(user).Return(1, nil)
			},

			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			id:          1,
			tokenBehavior: func(r *mock_service.MockAuthorization, token string) {

				r.EXPECT().ParseToken(token).Return(1, errors.New("invalid token"))
				//r.EXPECT().CheckStatus(1).Return(false, errors.New("something went wrong"))
			},

			inputBrandBody: `{"title":"Adidas", "description":"brand"}`,
			inputBrand:     marketplace.BrandsList{Title: "Adidas", Description: "brand"},
			mockBehavior: func(r *mock_service.MockBrands, brand marketplace.BrandsList) {

			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid token"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Init Dependences
			c := gomock.NewController(t)
			defer c.Finish()

			repoAuth := mock_service.NewMockAuthorization(c)
			test.authBehavior(repoAuth, test.inputUser)
			test.tokenBehavior(repoAuth, test.token)

			repoBrand := mock_service.NewMockBrands(c)
			test.mockBehavior(repoBrand, test.inputBrand)

			services := &service.Service{Authorization: repoAuth, Brands: repoBrand}
			handler := Handler{services}

			//Init routes
			r := gin.New()
			r.POST("/sign-up", handler.signUp)
			r.POST("/sign-in", handler.signIn)
			r.POST("/brands", handler.userIdentity, func(ctx *gin.Context) {
				status, _ := ctx.Get(userStatus)
				ctx.String(200, "%s", status)
			}, handler.createBrand)

			//Create request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(test.inputAuthBody))

			// Make Request
			r.ServeHTTP(w, req)

			w = httptest.NewRecorder()
			req = httptest.NewRequest("POST", "/sign-in",
				bytes.NewBufferString(test.inputAuthBody))

			r.ServeHTTP(w, req)

			w = httptest.NewRecorder()
			req = httptest.NewRequest("POST", "/brands",
				bytes.NewBufferString(test.inputBrandBody))
			req.Header.Set(test.headerName, test.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

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
