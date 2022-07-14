package app

import (
	"fmt"
	"go-assessment/internal/config"
	"go-assessment/internal/userrepository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_GetUser(t *testing.T) {
	testCases := []struct {
		name  string
		setup func(*Mockdb)

		expectedStatus int
		expectedBody   string
	}{{
		name: "Error",
		setup: func(md *Mockdb) {
			md.EXPECT().GetUser("1").Return(userrepository.User{}, fmt.Errorf("test error")).Times(1)
		},

		expectedStatus: http.StatusInternalServerError,
		expectedBody:   `{"error":"error getting user: test error","message":""}{"data":[{"id":"","first_name":"","last_name":""}]}`,
	}, {
		name: "Success",
		setup: func(md *Mockdb) {
			md.EXPECT().GetUser("1").
				Return(userrepository.User{Id: "1", FirstName: "test", LastName: "name"}, nil).
				Times(1)
		},

		expectedStatus: http.StatusOK,
		expectedBody:   `{"data":[{"id":"1","first_name":"test","last_name":"name"}]}`,
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockDatabase := NewMockdb(ctrl)
			tc.setup(mockDatabase)

			app := New(config.Config{}, mockDatabase)
			router := app.router()

			rw := httptest.NewRecorder()
			router.ServeHTTP(rw, httptest.NewRequest(http.MethodGet, "/user/1", nil))

			assert.Equal(t, tc.expectedStatus, rw.Code)
			assert.Equal(t, tc.expectedBody, rw.Body.String())
		})
	}
}
