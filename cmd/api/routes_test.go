package main

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DataDavD/snippetbox/greenlight/internal/data"
	"github.com/DataDavD/snippetbox/greenlight/internal/jsonlog"
	"github.com/DataDavD/snippetbox/greenlight/internal/mailer"
)

func newTestApplication() *application {
	cfg := config{
		port: 4000,
		env:  "test",
		db: struct {
			dsn          string
			maxOpenConns int
			maxIdleConns int
			maxIdleTime  string
		}{
			dsn:          "postgres://postgres:1245@localhost/mock?sslmode=disable",
			maxOpenConns: 10,
			maxIdleConns: 5,
			maxIdleTime:  "15m",
		},
		limiter: struct {
			rps     float64
			burst   int
			enabled bool
		}{
			rps:     2,
			burst:   4,
			enabled: true,
		},
		smtp: struct {
			host     string
			port     int
			username string
			password string
			sender   string
		}{
			host:     "smtp.mailtrap.io",
			port:     2525,
			username: "test_username",
			password: "test_password",
			sender:   "DoNotReply <test_sender@mailtrap.io>",
		},
		cors: struct {
			trustedOrigins []string
		}{
			trustedOrigins: []string{"http://localhost:4000"},
		},
	}

	logger := jsonlog.NewLogger(os.Stdout, jsonlog.LevelInfo)

	db, _ := sql.Open("postgres", "postgres://greenlight:1245@localhost/greenlight?sslmode=disable")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	return app
}

func TestHealthcheckHandler(t *testing.T) {
	// Create a new instance of your application
	app := newTestApplication()

	// Create an HTTP request to the healthcheck endpoint
	req, err := http.NewRequest(http.MethodGet, "http://localhost:4000/v1/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a recorder to record the HTTP response
	rr := httptest.NewRecorder()

	// Call the handler function to handle the request
	handler := app.routes()
	handler.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestListMoviesHandler(t *testing.T) {
	// Create a new instance of your application
	app := newTestApplication()

	// Create an HTTP request to list movies endpoint
	req, err := http.NewRequest(http.MethodGet, "/v1/movies", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer ZLMQP5PLMJEVXKNYWM4NTIHW2A")

	// Create a recorder to record the HTTP response
	rr := httptest.NewRecorder()

	// Call the handler function to handle the request
	handler := app.routes()
	handler.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateMovieHandler(t *testing.T) {
	// Create a new instance of your application
	app := newTestApplication()

	// Create a JSON payload for creating a movie
	payload := []byte(`{"title":"Test Movie","year":2024,"runtime":"120 mins","genres":["Action","Adventure"]}`)

	// Create an HTTP request with the payload
	req, err := http.NewRequest(http.MethodPost, "/v1/movies", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer ZLMQP5PLMJEVXKNYWM4NTIHW2A")
	// Create a recorder to record the HTTP response
	rr := httptest.NewRecorder()

	// Call the handler function to handle the request
	handler := app.routes()
	handler.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestShowMovieHandler(t *testing.T) {
	// Create a new instance of your application
	app := newTestApplication()

	// Create an HTTP request to fetch a movie by ID
	req, err := http.NewRequest(http.MethodGet, "/v1/movies/2", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer ZLMQP5PLMJEVXKNYWM4NTIHW2A")
	// Create a recorder to record the HTTP response
	rr := httptest.NewRecorder()

	// Call the handler function to handle the request
	handler := app.routes()
	handler.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
