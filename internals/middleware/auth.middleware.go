package middleware

import (
	"context"
	"errors"
	"net/http"
	"rest/api/internals/config"
	db "rest/api/internals/db/sqlc"
	"rest/api/internals/logger"
	"rest/api/internals/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Middleware struct {
	Store  db.Store
	Config *config.AppConfig
	Logger *logger.Logger
	Auth   utils.Auth
}

func NewMiddleware(params Middleware) *Middleware {
	return &Middleware{
		Store:  params.Store,
		Config: params.Config,
		Logger: params.Logger,
	}
}

type ContextKey string

const UserKey ContextKey = "user"

// a chi middleware
func (m *Middleware) AuthorizeUser() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// get token from the cookie
			token, err := r.Cookie("Authorization")
			if err != nil {
				m.Logger.Error("reading token from cookie: %v", err)
				utils.ErrorMessage(w, http.StatusUnauthorized, errors.New("unauthorized user"))
				return
			}

			// decode token
			decodedToken, err := m.Auth.DecodeToken(token.Value, m.Config.JwtSecret)
			if err != nil {
				m.Logger.Error("decoding token: %v", err)
				utils.InternalError(w, errors.New("something went wrong"))
				return
			}

			if claims, ok := decodedToken.Claims.(jwt.MapClaims); ok && decodedToken.Valid {
				// check expiry date
				if float64(time.Now().Unix()) > claims["exp"].(float64) {
					m.Logger.Error("decoded token expiry check")
					utils.ErrorMessage(w, http.StatusUnauthorized, errors.New("jwt expired"))
					return
				}
				// fetch the user with the ID 
				parsedUUID, _ := uuid.Parse(claims["user_id"].(string))
				// if err != nil {
				// 	m.Logger.Error("[uuid.Parse:] %v", err)
				// 	return errors.New("invalid ID format")
				// }

				user, err := m.Store.GetUser(r.Context(), pgtype.UUID{
					Bytes: parsedUUID,
					Valid: true,
				})
				if err != nil {
					m.Logger.Error("fetching user: %v", err)
					utils.InternalError(w, errors.New("something went wrong"))
					return
				}
				
				// add the user object to the request context
				ctx := context.WithValue(r.Context(), UserKey, user)

				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				m.Logger.Error("decodedToken claims failed")
				utils.ErrorMessage(w, http.StatusUnauthorized, errors.New("user not authorized"))
				return
			}
		})
	}
}
