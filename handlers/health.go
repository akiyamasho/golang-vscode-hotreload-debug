package handlers

import (
	"encoding/json"
	"net/http"
)

const STATUS_OK = "ok"
const HEADER_CONTENT_TYPE = "Content-Type"
const CONTENT_TYPE_JSON = "application/json"

type GetHealthResponse struct {
	Status string `json:"status"`
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	response := GetHealthResponse{Status: STATUS_OK}

	w.Header().Set(HEADER_CONTENT_TYPE, CONTENT_TYPE_JSON)
	json.NewEncoder(w).Encode(response)
}
