package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/knadh/koanf/v2"
	"github.com/orgs/the-tikka-cafe/tyme/internal/store"
	"github.com/orgs/the-tikka-cafe/tyme/internal/store/postgres"
	"github.com/zerodha/logf"
)

type App struct {
	store store.Store
	lo    logf.Logger
}

var (
	ko = koanf.New(".")

	buildString = "v0.0.1"
)

func main() {
	initConfig()

	app := &App{lo: initLogger(ko.Bool("app.enable_debug_logs"))}

	var conf postgres.Conf
	ko.UnmarshalWithConf("store.postgres", &conf, koanf.UnmarshalConf{Tag: "json"})

	postgres := postgres.New(conf)
	if postgres == nil {
		app.lo.Info("Failed to connect to Postgres")
	}

	app.store = postgres

	// Register Routers
	r := chi.NewRouter()
	r.Get("/", wrap(app, testRoute))

	srv := &http.Server{
		Addr:         ko.MustString("app.address"),
		ReadTimeout:  ko.MustDuration("app.server_timeout"),
		WriteTimeout: ko.MustDuration("app.server_timeout"),
		Handler:      r,
	}

	app.lo.Info("starting server", "address", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		app.lo.Fatal("couldn't start server", "error", err)
	}
}
