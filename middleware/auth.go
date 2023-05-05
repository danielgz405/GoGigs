package middleware

import (
	"net/http"
	"strings"

	"github.com/dg/acordia/models"
	"github.com/dg/acordia/repository"
	"github.com/dg/acordia/responses"
	"github.com/dg/acordia/server"
	"github.com/golang-jwt/jwt"
)

var (
	NO_AUTH_NEEDED = []string{
		"/welcome",
		"login",
		"signup",
		"user",
		"automobile",
		"visitor",
		"package",
		"events",
	}
)

func shouldCheckAuth(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

func CheckAuthMiddleware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckAuth(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
			_, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})
			if err != nil {
				responses.NoAuthResponse(w, http.StatusUnauthorized, "Expired or invalid token")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func ValidateToken(s server.Server, w http.ResponseWriter, r *http.Request) (*models.Profile, error) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})
	if err != nil {
		responses.NoAuthResponse(w, http.StatusUnauthorized, "Error validating token")
		return nil, err
	}
	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		userId := claims.UserId.Hex()
		profile, err := repository.GetUserById(r.Context(), userId)
		if err != nil {
			responses.NoAuthResponse(w, http.StatusUnauthorized, "Error validating token")
			return nil, err
		}
		return profile, nil
	} else {
		return nil, err
	}
}
