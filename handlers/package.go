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

type InsertPackageRequest struct {
	CompanyName        string `json:"company-name" bson:"company-name"`
	MessengerBoy       string `json:"messenger-boy" bson:"messenger-boy"`
	TypeIdentification string `json:"type-identification" bson:"type-identification"`
	Identification     string `json:"identification" bson:"identification"`
	Receiver           string `json:"receiver" bson:"receiver"`
	Tower              string `json:"tower" bson:"tower"`
	Office             string `json:"office" bson:"office"`
	Apto               string `json:"apto" bson:"apto"`
	PlaceSave          string `json:"place-save" bson:"place-save"`
	Autorization       bool   `json:"autorization" bson:"autorization"`
	Worker             string `json:"worker" bson:"worker"`
	Description        string `json:"description" bson:"description"`
}

func InsertPackageHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		var req = InsertPackageRequest{}
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

		packageEvent := models.InsertPackage{
			CompanyName:        req.CompanyName,
			MessengerBoy:       req.MessengerBoy,
			TypeIdentification: req.TypeIdentification,
			Identification:     req.Identification,
			Receiver:           req.Receiver,
			Tower:              req.Tower,
			Office:             req.Office,
			Apto:               req.Apto,
			PlaceSave:          req.PlaceSave,
			EventType:          "paquete",
			Date:               time.Now().In(loc).Format("2006-01-02 15:04:05"),
			Worker:             req.Worker,
			Description:        req.Description,
		}

		createdPackage, err := repository.InsertPackage(r.Context(), &packageEvent)
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdPackage)
	}
}
