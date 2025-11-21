package main

import (
	"net/http"
	"log"
	"time"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/vargascardona/neo-ledger/internal/guitars"
	repo "github.com/vargascardona/neo-ledger/internal/adapters/postgresql/sqlc"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("I'm good"))
	})

	guitarService := guitars.NewService(repo.New(app.db), app.db)
	guitarHandler := guitars.NewHandler(guitarService)
	r.Get("/guitars", guitarHandler.ListGuitars)
	r.Get("/guitars/{id}", guitarHandler.FindGuitarByID)
	r.Post("/guitars", guitarHandler.CreateGuitar)
	r.Patch("/guitars/{id}", guitarHandler.UpdateGuitar)
	r.Delete("/guitars/{id}", guitarHandler.DeleteGuitar)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("server has started at addr %s", app.config.addr)

	return srv.ListenAndServe()
}

type application struct {
	config config
	db *pgx.Conn
}

type config struct {
	addr string
	db dbConfig
}

type dbConfig struct {
	dsn string
}
