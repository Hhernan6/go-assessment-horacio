package app

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	internalConfig "go-assessment/internal/config"
	"go-assessment/internal/userrepository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Start initializes and runs the webserver
// Copied from https://github.com/gorilla/mux#graceful-shutdown
func Start() {
	var wait time.Duration
	var configFilePath string
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&configFilePath, "config-filepath", "../../config/config.json", "the file path for the config file")
	flag.Parse()

	// attempt to load config from config file path
	cfg, err := internalConfig.GetConfig(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	if nil != err {
		log.Fatalln(err)
	}

	ctx := context.Background()

	dbEndpoint := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	fmt.Println(dbEndpoint)
	db, err := sql.Open("mysql", dbEndpoint)
	if err != nil {
		panic(err)
	}

	userRepository := userrepository.New(db)

	application := New(cfg, userRepository)

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", cfg.Application.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      application.router(), // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
