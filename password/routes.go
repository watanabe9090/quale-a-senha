package password

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/watanabe9090/quale-a-senha/internal"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/new-password", handleGeneratePassword)
}

func handleGeneratePassword(w http.ResponseWriter, r *http.Request) {
	var length = 16
	var err error
	if r.URL.Query().Get("length") != "" {
		length, err = strconv.Atoi(r.URL.Query().Get("length"))
		if err != nil {
			w.Header().Add("Content-type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(internal.APIResponse{
				Message: err.Error(),
				Data:    nil,
			})
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(internal.APIResponse{
		Message: "OK",
		Data:    GeneratePassword(length),
	})
}
