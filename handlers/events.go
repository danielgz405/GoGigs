package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dg/acordia/models"
	"github.com/dg/acordia/repository"
	"github.com/dg/acordia/responses"
	"github.com/dg/acordia/server"
)

func ListEventsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		event, err := repository.ListEvents(r.Context())
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		if event == nil {
			event = []models.Event{}
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(event)
	}
}
