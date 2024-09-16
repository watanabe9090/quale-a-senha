package password

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/watanabe9090/quale-a-senha/internal"
)

var (
	mux     *http.ServeMux
	respRec *httptest.ResponseRecorder
)

func setup() {
	mux = http.NewServeMux()
	RegisterRoutes(mux)
	respRec = httptest.NewRecorder()
}

func Test_Routes(t *testing.T) {
	setup()
	req, err := http.NewRequest("GET", "/api/v1/new-password", nil)
	if err != nil {
		t.Fatal("Creating 'GET /api/v1/new-password' request failed!")
	}
	mux.ServeHTTP(respRec, req)
	if respRec.Result().StatusCode != 200 {
		log.Fatalf("statusCode %d should be 200", respRec.Result().StatusCode)
	}
	if respRec.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("Content-type %s should be application/json", respRec.Header().Get("Content-Type"))
	}
}

func Test_Routes2(t *testing.T) {
	setup()
	req, err := http.NewRequest("GET", "/api/v1/new-password", nil)
	if err != nil {
		t.Fatal("Creating 'GET /api/v1/new-password' request failed!")
	}
	mux.ServeHTTP(respRec, req)
	if respRec.Result().StatusCode != 200 {
		log.Fatalf("statusCode %d should be 200", respRec.Result().StatusCode)
	}
}

func Test_Routes3(t *testing.T) {
	setup()
	req, err := http.NewRequest("GET", "/api/v1/new-password?length=32", nil)
	if err != nil {
		t.Fatal("Creating 'GET /api/v1/new-password' request failed!")
	}
	mux.ServeHTTP(respRec, req)
	if respRec.Result().StatusCode != 200 {
		log.Fatalf("statusCode %d should be 200", respRec.Result().StatusCode)
	}
	var body internal.APIResponse
	json.NewDecoder(respRec.Body).Decode(&body)

	if len(body.Data.(string)) != 32 {
		log.Fatalf("generated password %d should have length 32", len(body.Data.(string)))
	}
}

func Test_Routes4(t *testing.T) {
	setup()
	req, err := http.NewRequest("GET", "/api/v1/new-password?length=abc", nil)
	if err != nil {
		t.Fatal("Creating 'GET /api/v1/new-password' request failed!")
	}
	mux.ServeHTTP(respRec, req)
	if respRec.Result().StatusCode != 400 {
		log.Fatalf("statusCode %d should be 400", respRec.Result().StatusCode)
	}
	var body internal.APIResponse
	json.NewDecoder(respRec.Body).Decode(&body)
}
