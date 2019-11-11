package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response_body, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response_body)
}

func RespondSuccess(w http.ResponseWriter) {
	respondWithJSON(w, http.StatusOK, map[string]bool{"success": true})
}

func RespondBadRequest(w http.ResponseWriter, errCode int, err error) {
	respondWithJSON(w, http.StatusBadRequest, ErrorResponse{errCode, err.Error()})
}

func RespondUnauthorized(w http.ResponseWriter, errCode int, err error) {
	respondWithJSON(w, http.StatusUnauthorized, ErrorResponse{errCode, err.Error()})
}
