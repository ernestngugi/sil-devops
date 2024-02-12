package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/ernestngugi/sil-devops/internal/db"
	"github.com/ernestngugi/sil-devops/internal/web/router.go"
)

const defaultPort = "3000"

func main() {

	dB := db.InitDB()
	defer dB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	appRouter := router.BuildRouter(dB)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: appRouter,
	}

	fmt.Printf("api starting, listening on :%v", port)

	if err := server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("server shutdown")
		} else {
			fmt.Printf("server shutdown unexpectedly %v", err)
		}
	}
}
