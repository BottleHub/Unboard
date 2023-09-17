package auth

import (
	"context"
	"net/http"

	"github.com/bottlehub/unboard/configs"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			// Validate jwt token
			tokenStr := header
			username, err := configs.ParseToken(tokenStr, secretKey)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// Create user and check if user exists in db
			user := configs.User{Username: username}
			id, err := configs.ConnectDB().GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = id
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// Call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// Finds the user from the context. REQUIRES Middleware to have run.
func FromContext(ctx context.Context) *configs.User {
	raw, _ := ctx.Value(userCtxKey).(*configs.User)
	return raw
}

func main() {
	
}