package main

import (
	"fmt"
	"net/http"

	"github.com/caosbad/ever-post-mixin-bot/config"
	"github.com/caosbad/ever-post-mixin-bot/durable"
	"github.com/caosbad/ever-post-mixin-bot/middlewares"
	"github.com/caosbad/ever-post-mixin-bot/routes"
	"github.com/dimfeld/httptreemux"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/go-pg/pg"
	"github.com/unrolled/render"
)

func StartServer(db *pg.DB) error {
	logger, err := durable.NewLoggerClient("", true)
	if err != nil {
		return err
	}
	defer logger.Close()

	router := httptreemux.New()
	routes.RegisterHanders(router)
	routes.RegisterRoutes(router)
	handler := middlewares.Authenticate(router)
	handler = middlewares.Constraint(handler)
	handler = middlewares.Context(handler, db, render.New(render.Options{UnEscapeHTML: true}))
	return gracehttp.Serve(&http.Server{Addr: fmt.Sprintf(":%d", config.HTTPListenPort), Handler: handler})
}
