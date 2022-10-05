package middleware

import (
	"context"
	"github.com/Julia1505/RedditCloneBack/pkg/user"
	"net/http"
	"strings"
)

func IsAuthorized(st user.UsersRepo, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		token := strings.Replace(bearer, "Bearer ", "", 1)
		if token != "" {
			user, err := st.GetByToken(token)
			if err != nil {
				http.Error(w, "no auth", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		next.ServeHTTP(w, r)
	})
}
