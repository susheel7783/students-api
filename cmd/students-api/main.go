package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/susheel7783/students-api/internal/config"
	"github.com/susheel7783/students-api/internal/http/handlers/student"
	"github.com/susheel7783/students-api/internal/storage/sqlite"
)

func main() {
	// fmt.Println("welcome to students api")
	// load config
	cfg := config.MustLoad()

	// database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))
	// setup routes
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))

	// get by id
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	// get list of students
	router.HandleFunc("GET /api/students", student.GetList(storage))

	// router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Welcome to students-api"))
	// })

	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server Started", slog.String("address", cfg.Addr))
	// fmt.Printf("Server started %s", cfg.Addr)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shut down server", slog.String("error", err.Error()))
	}

	slog.Info("server shut down gracefully")

}

// to run this file, use the command: go run cmd/students-api/main.go
// to run the server we have to create a config file in yaml format and give the path of that file in environment variable CONFIG_PATH
// go run cmd/students-api/main.go -config config/local.yaml
// and we can see the students table cretae in the storage.db file in the root directory
