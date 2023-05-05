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

type InsertVisitorRequest struct {
	Into               bool   `json:"into" bson:"into"`
	TypeIdentification string `json:"type-identification" bson:"type-identification"`
	Identification     string `json:"identification" bson:"identification"`
	Tower              string `json:"tower" bson:"tower"`
	Floor              string `json:"floor" bson:"floor"`
	Apto               string `json:"apto" bson:"apto"`
	Automobile         bool   `json:"automobile" bson:"automobile"`
	Autorization       bool   `json:"autorization" bson:"autorization"`
	Worker             string `json:"worker" bson:"worker"`
	Description        string `json:"description" bson:"description"`
}

func InsertVisitorHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		var req = InsertVisitorRequest{}
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

		visitor := models.InsertVisitor{
			Into:               req.Into,
			TypeIdentification: req.TypeIdentification,
			Identification:     req.Identification,
			Tower:              req.Tower,
			Floor:              req.Floor,
			Apto:               req.Apto,
			Automobile:         req.Automobile,
			EventType:          "ingreso",
			Date:               time.Now().In(loc).Format("2006-01-02 15:04:05"),
			Worker:             req.Worker,
			Description:        req.Description,
		}

		createdVisitor, err := repository.InsertVisitor(r.Context(), &visitor)
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdVisitor)
	}
}
