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

func TestHandler_signUp(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_service.MockAuthorization, user marketplace.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            marketplace.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"username": "username", "name": "name", "password": "password"}`,
			inputUser: marketplace.User{
				Username: "username",
				Name:     "name",
				Password: "password",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user marketplace.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"username": "username"}`,
			inputUser:            marketplace.User{},
			mockBehavior:         func(r *mock_service.MockAuthorization, user marketplace.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"username": "username", "name": "name", "password": "password"}`,
			inputUser: marketplace.User{
				Username: "username",
				Name:     "name",
				Password: "password",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user marketplace.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("something went wrong"))
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

			repo := mock_service.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_signIn(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAuthorization, user marketplace.User)

	tests := []struct {
		name               string
		inputBody          string
		inputUser          marketplace.User
		mockBehavior       mockBehavior
		expectedStatusCode int
		//expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"username": "username", "name": "name", "password": "password"}`,
			inputUser: marketplace.User{
				Username: "username",
				Name:     "name",
				Password: "password",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user marketplace.User) {
				r.EXPECT().GenerateToken(user.Username, user.Password).Return("64666873727468626378767a6668735baa61e4c9b93f3f0682250b6cf8331b7ee68fd8", nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:               "Wrong Input",
			inputBody:          `{"username": "username"}`,
			inputUser:          marketplace.User{},
			mockBehavior:       func(r *mock_service.MockAuthorization, user marketplace.User) {},
			expectedStatusCode: 400,
		},
		{
			name:      "Service Error",
			inputBody: `{"username": "username", "name": "name", "password": "password"}`,
			inputUser: marketplace.User{
				Username: "username",
				Name:     "name",
				Password: "password",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user marketplace.User) {
				r.EXPECT().GenerateToken(user.Username, user.Password).Return("", errors.New("something went wrong"))
			},
			expectedStatusCode: 500,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(repo, testCase.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/sign-in", handler.signIn)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in",
				bytes.NewBufferString(testCase.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
		})
	}
}
