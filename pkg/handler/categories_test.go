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

func TestHandler_getImage(t *testing.T) {
	type mockBehavior func(r *mock_service.MockCategories, id int)

	tests := []struct {
		name string

		idCategory           int
		inputBodyUri         string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:         "Invalid id",
			idCategory:   1,
			inputBodyUri: "/brands/image/p",
			mockBehavior: func(r *mock_service.MockCategories, id int) {
				//r.EXPECT().GetImage(id).Return("./static/categories/1/1.jpeg", nil)
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid id param"}`,
		},
		{
			name:         "Service Error",
			idCategory:   1,
			inputBodyUri: "/brands/image/1",
			mockBehavior: func(r *mock_service.MockCategories, id int) {
				r.EXPECT().GetImage(id).Return("", errors.New("something went wrong"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoCategory := mock_service.NewMockCategories(c)
			test.mockBehavior(repoCategory, test.idCategory)

			services := &service.Service{Categories: repoCategory}
			handler := Handler{services}

			r := gin.New()
			r.GET("/brands/image/:id", handler.getCategoriesImage)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", test.inputBodyUri, nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_deleteCategories(t *testing.T) {
	type mockBehavior func(r *mock_service.MockCategories, id int)
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

		idCategories int
		mockBehavior mockBehavior

		expectedStatusCode   int
		expectedResponseBody string
		uri                  string
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

			idCategories: 1,
			mockBehavior: func(r *mock_service.MockCategories, id int) {
				r.EXPECT().Delete(1).Return(nil)
			},
			uri:                  "/categories/1",
			expectedStatusCode:   200,
			expectedResponseBody: `%!s(bool=true){"status":"ok"}`,
		},
		{
			name:          "Have not rights",
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

				r.EXPECT().ParseToken(token).Return(1, nil)
				r.EXPECT().CheckStatus(1).Return(false, nil)
			},

			idCategories: 1,
			mockBehavior: func(r *mock_service.MockCategories, id int) {

			},
			uri:                  "/categories/1",
			expectedStatusCode:   200,
			expectedResponseBody: `%!s(bool=false){"message":"you do not have sufficient rights"}`,
		},
		{
			name:          "Invalid id param",
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

			idCategories: 1,
			mockBehavior: func(r *mock_service.MockCategories, id int) {

			},
			uri:                  "/categories/p",
			expectedStatusCode:   200,
			expectedResponseBody: `%!s(bool=true){"message":"invalid id param"}`,
		},
		{
			name:          "Server service error",
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

			idCategories: 1,
			mockBehavior: func(r *mock_service.MockCategories, id int) {
				r.EXPECT().Delete(1).Return(errors.New(""))
			},
			uri:                  "/categories/1",
			expectedStatusCode:   200,
			expectedResponseBody: `%!s(bool=true){"message":"this brand not exists"}`,
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

			repoCategory := mock_service.NewMockCategories(c)
			test.mockBehavior(repoCategory, test.idCategories)

			services := &service.Service{Authorization: repoAuth, Categories: repoCategory}
			handler := Handler{services}

			//Init routes
			r := gin.New()
			r.POST("/sign-up", handler.signUp)
			r.POST("/sign-in", handler.signIn)
			r.POST("/categories/:id", handler.userIdentity, func(ctx *gin.Context) {
				status, _ := ctx.Get(userStatus)
				ctx.String(200, "%s", status)
			}, handler.deleteCategory)

			//Create request
			w1 := httptest.NewRecorder()
			req1 := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(test.inputAuthBody))

			// Make Request
			r.ServeHTTP(w1, req1)

			w2 := httptest.NewRecorder()
			req2 := httptest.NewRequest("POST", "/sign-in",
				bytes.NewBufferString(test.inputAuthBody))

			r.ServeHTTP(w2, req2)

			w3 := httptest.NewRecorder()
			req3 := httptest.NewRequest("POST", test.uri, nil)
			req3.Header.Set(test.headerName, test.headerValue)

			r.ServeHTTP(w3, req3)

			assert.Equal(t, w3.Code, test.expectedStatusCode)
			assert.Equal(t, w3.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHanler_createCategory(t *testing.T) {
	type mockBehavior func(r *mock_service.MockCategories, category marketplace.CategoriesList)
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

		inputCategoryBody string
		inputCategories   marketplace.CategoriesList
		mockBehavior      mockBehavior

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

			inputCategoryBody: `{"title":"Shoes"}`,
			inputCategories:   marketplace.CategoriesList{Title: "Shoes"},
			mockBehavior: func(r *mock_service.MockCategories, category marketplace.CategoriesList) {
				r.EXPECT().GetByString(category.Title).Return(1, errors.New(""))
				r.EXPECT().Create(category).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `%!s(bool=true){"id":1}`,
		},
		{
			name:          "Invalid category",
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

			inputCategoryBody: `{"description":"brand"}`,
			inputCategories:   marketplace.CategoriesList{Title: "brand"},
			mockBehavior: func(r *mock_service.MockCategories, category marketplace.CategoriesList) {
			},
			expectedStatusCode:   200,
			expectedResponseBody: `%!s(bool=true){"message":"invalid param"}`,
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

			inputCategoryBody: `{"title":"Shoes"}`,
			inputCategories:   marketplace.CategoriesList{Title: "Shoes"},
			mockBehavior: func(r *mock_service.MockCategories, category marketplace.CategoriesList) {
				r.EXPECT().GetByString(category.Title).Return(1, nil)
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

			inputCategoryBody: `{"title":"Shoes"}`,
			inputCategories:   marketplace.CategoriesList{Title: "Shoes"},
			mockBehavior: func(r *mock_service.MockCategories, category marketplace.CategoriesList) {
				r.EXPECT().GetByString(category.Title).Return(1, errors.New(""))
				r.EXPECT().Create(category).Return(0, errors.New("something went wrong"))
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

			inputCategoryBody: `{"title":"Shoes"}`,
			inputCategories:   marketplace.CategoriesList{Title: "Shoes"},
			mockBehavior: func(r *mock_service.MockCategories, category marketplace.CategoriesList) {

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

			inputCategoryBody: `{"title":"Shoes"}`,
			inputCategories:   marketplace.CategoriesList{Title: "Shoes"},
			mockBehavior: func(r *mock_service.MockCategories, category marketplace.CategoriesList) {

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

			inputCategoryBody: `{"title":"Shoes"}`,
			inputCategories:   marketplace.CategoriesList{Title: "Shoes"},
			mockBehavior: func(r *mock_service.MockCategories, category marketplace.CategoriesList) {

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

			repoCategory := mock_service.NewMockCategories(c)
			test.mockBehavior(repoCategory, test.inputCategories)

			services := &service.Service{Authorization: repoAuth, Categories: repoCategory}
			handler := Handler{services}

			//Init routes
			r := gin.New()
			r.POST("/sign-up", handler.signUp)
			r.POST("/sign-in", handler.signIn)
			r.POST("/brands", handler.userIdentity, func(ctx *gin.Context) {
				status, _ := ctx.Get(userStatus)
				ctx.String(200, "%s", status)
			}, handler.createCategories)

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
				bytes.NewBufferString(test.inputCategoryBody))
			req.Header.Set(test.headerName, test.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

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
