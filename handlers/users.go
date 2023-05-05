package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dg/acordia/middleware"
	"github.com/dg/acordia/models"
	"github.com/dg/acordia/repository"
	"github.com/dg/acordia/responses"
	"github.com/dg/acordia/server"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type SignUpLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UpdateUserRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Image     string `json:"image"`
	DesertRef string `json:"desertref"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var req = SignUpLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		createUser := models.InsertUser{
			Email:    req.Email,
			Password: string(hashedPassword),
			Name:     req.Name,
		}
		profile, err := repository.InsertUser(r.Context(), &createUser)
		if err != nil {
			responses.BadRequest(w, "Error creating user")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(profile)
	}
}

func LoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req = SignUpLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		user, err := repository.GetUserByEmail(r.Context(), req.Email)
		if user == nil {
			responses.NoAuthResponse(w, http.StatusUnauthorized, "Invalid credentials")
			return
		}
		var profile = models.Profile{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			Image:     user.Image,
			DesertRef: user.DesertRef,
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			responses.NoAuthResponse(w, http.StatusUnauthorized, "Invalid credentials")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(profile)
	}
}

func ProfileHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		profile, err := repository.GetUserById(r.Context(), params["userId"])
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}
		// Handle request
		json.NewEncoder(w).Encode(profile)
	}
}

func UpdateUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		user, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		var req = UpdateUserRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		data := models.UpdateUser{
			Id:        user.Id.Hex(),
			Name:      req.Name,
			Email:     req.Email,
			Image:     req.Image,
			DesertRef: req.DesertRef,
		}
		updatedUser, err := repository.UpdateUser(r.Context(), data)
		if err != nil {
			responses.BadRequest(w, "Error updating user")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedUser)
	}
}

func DeleteUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		user, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		err = repository.DeleteUser(r.Context(), user.Id.Hex())
		if err != nil {
			responses.BadRequest(w, "Error deleting user")
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
