package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-api/common/jwt"
	"rest-api/common/middleware"
	"rest-api/common/routes"
	"rest-api/features/post"
	"rest-api/features/secret"
	"time"

	"github.com/go-playground/validator/v10"
)

func NewServer(l *log.Logger) http.Handler {
	mux := http.NewServeMux()

	v := validator.New(validator.WithRequiredStructEnabled())

	postHandlers := post.NewPostHandlers(post.NewPostRepository(), v, l)

	jwtSecret := "willbefixedlaterignorefornow"
	jwtService := jwt.NewJwtService(jwtSecret, l)
	secretHandlers := secret.NewSecretHandlers(l)

	handlers := routes.Handlers{
		Post:   postHandlers,
		Secret: secretHandlers,
	}

	routes.SetupRoutes(mux, handlers, l, jwtService)

	logMiddleware := middleware.LogMiddleware(l)
	middlewares := []middleware.Middleware{
		logMiddleware,
		middleware.JsonMiddleware,
	}

	globalMiddleware := middleware.NewChain(middlewares...)

	return globalMiddleware.Apply(mux)
}

func run(ctx context.Context, args []string, getenv func(string) string, stdout, stderr io.Writer) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	l := log.New(stdout, "LOG: ", log.LstdFlags|log.Lshortfile)

	mux := NewServer(l)

	srv := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	go func() {
		l.Printf("Server starting on http://localhost%s", srv.Addr)

		err := srv.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(stderr, "Server failed: %v\n", err)
		}
	}()

	<-ctx.Done()

	l.Println("Received shutdown signal. Preparing to shut down gracefully...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	err := srv.Shutdown(shutdownCtx)

	if err != nil {
		fmt.Fprintf(stderr, "Server forced to shutdown: %v\n", err)
		return err
	}

	l.Println("Server gracefully stopped. Goodbye!")
	return nil
}

func main() {
	err := run(context.Background(), os.Args, os.Getenv, os.Stdout, os.Stderr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Application exited with error: %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
