package response

import (
	"encoding/json"
	"net/http"

	ulanderrors "userland/errors"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	responseBody, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseBody)
}

func RespondSuccess(w http.ResponseWriter) {
	respondWithJSON(w, http.StatusOK, map[string]bool{"success": true})
}

func RespondSuccessWithBody(w http.ResponseWriter, payload interface{}) {
	respondWithJSON(w, http.StatusOK, payload)
}

func RespondBadRequest(w http.ResponseWriter, err ulanderrors.UserlandError) {
	respondWithJSON(w, http.StatusBadRequest, err)
}

func RespondUnauthorized(w http.ResponseWriter, err ulanderrors.UserlandError) {
	respondWithJSON(w, http.StatusUnauthorized, err)
}

func RespondInternalError(w http.ResponseWriter, err ulanderrors.UserlandError) {
	respondWithJSON(w, http.StatusInternalServerError, err)
}
