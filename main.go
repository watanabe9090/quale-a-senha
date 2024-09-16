package main

import (
	"net/http"

	"github.com/watanabe9090/quale-a-senha/password"
)

func main() {
	mux := http.NewServeMux()
	password.RegisterRoutes(mux)

	http.ListenAndServe(":3001", mux)
}
