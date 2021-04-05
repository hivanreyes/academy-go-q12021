package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hivanreyes/academy-go-q12021/router/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//mockgen -source=router/routes.go -destination=router/mocks/routes.go -package=mocks

func TestNew(t *testing.T) {
	testCases := []struct {
		name           string
		endpoint       string
		handlerName    string
		status         int
		callController bool
	}{
		{
			name:           "Ok, Get All Pokemons",
			endpoint:       "/getAllPokemon",
			handlerName:    "getAll",
			status:         200,
			callController: true,
		},
		{
			name:        "Ok, Populate Pokemons",
			endpoint:    "/populateAllPokemon",
			handlerName: "populateAll",
			status:      200,
		},
		{
			name:        "Ok, Get pokemons concurrently",
			endpoint:    "/getConcurrentPokemon",
			handlerName: "concurrent",
			status:      200,
		},
		{
			name:        "Not Found, Calling endpoint does not exist return 404",
			endpoint:    "/doesNotExistEndpoint",
			handlerName: "doesnotexist",
			status:      404,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			c := mocks.NewMockController(mockCtrl)

			r := New(c)

			if tc.handlerName == "getAll" {
				c.EXPECT().ReadPokemon(gomock.Any(), gomock.Any()).Times(1)
			} else if tc.handlerName == "populateAll" {
				c.EXPECT().SavePokemon(gomock.Any(), gomock.Any()).Times(1)
			} else if tc.handlerName == "concurrent" {
				c.EXPECT().ReadConcurrentPokemon(gomock.Any(), gomock.Any()).Times(1)
			}

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, tc.endpoint, nil)

			r.ServeHTTP(recorder, request)
			assert.Equal(t, tc.status, recorder.Code)
			assert.Nil(t, err)
		})
	}
}
