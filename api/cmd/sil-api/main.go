package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/ernestngugi/sil-devops/internal/db"
	"github.com/ernestngugi/sil-devops/internal/web/router.go"
	"github.com/joho/godotenv"
)

const defaultPort = "3000"

func main() {

	var envFilePath string
	flag.StringVar(&envFilePath, "e", "", "Path to env file")
	flag.Parse()

	if envFilePath != "" {
		err := godotenv.Load(envFilePath)
		if err != nil {
			panic(fmt.Errorf("failed to load env file: %v", err))
		}
	}

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
