package middlewares

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"github.com/caosbad/ever-post-mixin-bot/models"
	"github.com/caosbad/ever-post-mixin-bot/session"
	"github.com/caosbad/ever-post-mixin-bot/views"
	// "github.com/caosbad/telegraph"
)

var whitelist = [][2]string{
	{"GET", "^/$"},
	{"POST", "^/auth$"},
	{"GET", "^/posts"},
	{"GET", "^/posts/.*"},
	{"GET", "^/user/.*"},
}

type contextValueKey struct{ int }

var keyCurrentUser = contextValueKey{1000}
// get cached user value 
func CurrentUser(r *http.Request) *models.User {
	user, _ := r.Context().Value(keyCurrentUser).(*models.User)
	return user
}

// auth an cache user in context
func Authenticate(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			handleUnauthorized(handler, w, r)
			return
		}

		user, err := models.AuthenticateUserByToken(r.Context(), header[7:])
		if err != nil {
			views.RenderErrorResponse(w, r, err)
		} else if user == nil {
			handleUnauthorized(handler, w, r)
		} else {
			ctx := context.WithValue(r.Context(), keyCurrentUser, user)
			handler.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func handleUnauthorized(handler http.Handler, w http.ResponseWriter, r *http.Request) {
	for _, pp := range whitelist {
		if pp[0] != r.Method {
			continue
		}
		if matched, err := regexp.MatchString(pp[1], strings.ToLower(r.URL.Path)); err == nil && matched {
			handler.ServeHTTP(w, r)
			return
		}
	}

	views.RenderErrorResponse(w, r, session.AuthorizationError(r.Context()))
}
