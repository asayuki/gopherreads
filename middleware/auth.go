package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/asayuki/gopherreads/config"
	"github.com/asayuki/gopherreads/stores"
	"github.com/golang-jwt/jwt"
)

type contextKey string

const UserKey contextKey = "sub"

func Auth(user *stores.UserStore) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("auth_token")
			if err != nil {
				http.Redirect(w, r, "/auth", http.StatusFound)
				return
			}

			tokenString := c.Value

			token, err := validateJWT(tokenString)
			if err != nil {
				http.Redirect(w, r, "/auth", http.StatusFound)
				return
			}

			claims := token.Claims.(jwt.MapClaims)
			str := claims["sub"].(float64)

			user, err := user.GetUserByField("id", str)
			if err != nil {
				c := &http.Cookie{
					Name:     "auth_token",
					Value:    "",
					Path:     "/",
					Expires:  time.Unix(0, 0),
					HttpOnly: true,
				}

				http.SetCookie(w, c)
				http.Redirect(w, r, "/auth", http.StatusFound)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, UserKey, user.ID)
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

func validateJWT(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.Envs.SessionSecret), nil
	})
}
