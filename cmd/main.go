package main

import (
	"log"
	"net/http"
	"simple-todo/config"
	"simple-todo/internal/adapter/http/handler"
	"simple-todo/internal/adapter/http/router"
	"simple-todo/internal/infrastructure/postgres/repository"
	"simple-todo/internal/usecase/service"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.Load()

	db, err := sqlx.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	m, err := migrate.New(
		"file://migrations",
		cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Could not connect to migrations: %v", err)
	}
	if err := m.Up(); err != nil {
		log.Printf("Could not run migrations: %v", err)
	}

	repo := repository.NewTodoRepository(db)
	service := service.NewTodoService(repo)
	handler := handler.NewTodoHandler(service)
	router := router.NewRouter(handler)

	log.Printf("Server started on 127.0.0.1:%s", cfg.Port)

	http.ListenAndServe(":"+cfg.Port, router)
}
