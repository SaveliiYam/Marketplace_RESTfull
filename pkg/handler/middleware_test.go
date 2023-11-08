package handler

import (
	"errors"
	"marketplace/pkg/service"
	mock_service "marketplace/pkg/service/mocks"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		id                   int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			id:          1,
			mockBehavior: func(r *mock_service.MockAuthorization, token string) {
				r.EXPECT().ParseToken(token).Return(1, nil)
				r.EXPECT().CheckStatus(1)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "1",
		},
		{
			name:                 "Invalid Header Name",
			headerName:           "",
			headerValue:          "Bearer token",
			token:                "token",
			mockBehavior:         func(r *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"empty auth header"}`,
		},
		{
			name:                 "Invalid Header Value",
			headerName:           "Authorization",
			headerValue:          "Bearr token",
			token:                "token",
			mockBehavior:         func(r *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name:                 "Empty Token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			token:                "token",
			mockBehavior:         func(r *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"token is empty"}`,
		},
		{
			name:        "Parse Error",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *mock_service.MockAuthorization, token string) {
				r.EXPECT().ParseToken(token).Return(0, errors.New("invalid token"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid token"}`,
		},
		{
			name:        "User status Error",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			id:          1,
			mockBehavior: func(r *mock_service.MockAuthorization, token string) {
				r.EXPECT().ParseToken(token).Return(1, nil)
				r.EXPECT().CheckStatus(1).Return(false, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			test.mockBehavior(repo, test.token)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.GET("/identity", handler.userIdentity, func(c *gin.Context) {
				id, _ := c.Get(userCtx)
				c.String(200, "%d", id)
			})

			// Init Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/identity", nil)
			req.Header.Set(test.headerName, test.headerValue)

			r.ServeHTTP(w, req)

			// Asserts
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestCheckStatus(t *testing.T) {
	var getContext = func(status bool) *gin.Context {
		stat := &gin.Context{}
		stat.Set(userStatus, status)
		return stat
	}
	testTable := []struct {
		name       string
		stat       *gin.Context
		status     bool
		shouldFail bool
	}{
		{
			name:   "OK",
			stat:   getContext(true),
			status: true,
		},
		{
			name:       "Empty",
			stat:       &gin.Context{},
			shouldFail: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			status, err := checkStatus(test.stat)
			if test.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, status, test.status)
		})
	}
}

func TestGetUserId(t *testing.T) {
	var getContext = func(id int) *gin.Context {
		ctx := &gin.Context{}
		ctx.Set(userCtx, id)
		return ctx
	}

	testTable := []struct {
		name       string
		ctx        *gin.Context
		id         int
		shouldFail bool
	}{
		{
			name: "Ok",
			ctx:  getContext(1),
			id:   1,
		},
		{
			name:       "Empty",
			ctx:        &gin.Context{},
			shouldFail: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			id, err := getUserId(test.ctx)
			if test.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, id, test.id)
		})
	}
}
