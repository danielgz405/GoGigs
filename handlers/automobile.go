package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dg/acordia/models"
	"github.com/dg/acordia/repository"
	"github.com/dg/acordia/responses"
	"github.com/dg/acordia/server"
)

type InsertAutomobileRequest struct {
	Into           bool   `json:"into" bson:"into"`
	CarType        string `json:"carType" bson:"carType"`
	Identification string `json:"identification" bson:"identification"`
	Autorization   bool   `json:"autorization" bson:"autorization"`
	Worker         string `json:"worker" bson:"worker"`
	Description    string `json:"description" bson:"description"`
}

func InsertAutomobileHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		var req = InsertAutomobileRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}
		loc, error := time.LoadLocation("America/Bogota")
		if error != nil {
			responses.InternalServerError(w, "Error loading location")
			return
		}

		automobile := models.InsertAutomobile{
			Into:           req.Into,
			CarType:        req.CarType,
			Identification: req.Identification,
			Autorization:   req.Autorization,
			EventType:      "vehiculo",
			Date:           time.Now().In(loc).Format("2006-01-02 15:04:05"),
			Worker:         req.Worker,
			Description:    req.Description,
		}

		createdAutomobile, err := repository.InsertAutomobile(r.Context(), &automobile)
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdAutomobile)
	}
}
