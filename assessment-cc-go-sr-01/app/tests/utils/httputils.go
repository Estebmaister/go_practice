package utilstests

import (
	"battle-of-monsters/app/router"
	"net/http"
	"net/http/httptest"
)

func ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	r := router.Router()
	nr := httptest.NewRecorder()

	r.ServeHTTP(nr, req)

	return nr
}
