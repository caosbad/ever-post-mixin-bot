package middlewares

import (
	"net/http"

	"github.com/caosbad/ever-post-mixin-bot/session"
	"github.com/go-pg/pg"
	"github.com/unrolled/render"
)

func Context(handler http.Handler, db *pg.DB, render *render.Render) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := session.WithRequest(r.Context(), r)
		ctx = session.WithDatabase(ctx, db)
		ctx = session.WithRender(ctx, render)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
