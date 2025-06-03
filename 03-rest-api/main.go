package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-api/common/middleware"
	"rest-api/common/routes"
	"rest-api/features/post"
	"rest-api/features/secret"
	"time"

	"github.com/go-playground/validator/v10"
)

func run(ctx context.Context, args []string, getenv func(string) string, stdout, stderr io.Writer) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	mux := http.NewServeMux()

	v := validator.New(validator.WithRequiredStructEnabled())
	l := log.New(stdout, "LOG: ", log.LstdFlags|log.Lshortfile)

	postHandlers := post.NewPostHandlers(post.NewPostRepository(), v, l)
	secretHandlers := secret.NewSecretHandlers(l)

	handlers := routes.Handlers{
		Post:   postHandlers,
		Secret: secretHandlers,
	}

	routes.SetupRoutes(mux, handlers, l)

	middlewares := []middleware.Middleware{
		middleware.LogMiddleware,
		middleware.JsonMiddleware,
	}
	globalMiddleware := middleware.NewChain(middlewares...)

	srv := &http.Server{
		Addr:    ":3000",
		Handler: globalMiddleware.Apply(mux),
	}

	go func() {
		l.Printf("Server starting on http://localhost%s", srv.Addr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(stderr, "Server failed: %v\n", err)
		}
	}()

	<-ctx.Done()

	l.Println("Received shutdown signal. Preparing to shut down gracefully...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		fmt.Fprintf(stderr, "Server forced to shutdown: %v\n", err)
		return err
	}

	l.Println("Server gracefully stopped. Goodbye!")
	return nil
}

func main() {
	appCtx := context.Background()

	if err := run(appCtx, os.Args, os.Getenv, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "Application exited with error: %s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
