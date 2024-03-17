package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sophie-rigg/csv-file-store/cache"
	"github.com/sophie-rigg/csv-file-store/server"
	"github.com/sophie-rigg/csv-file-store/storage"
	"github.com/sophie-rigg/csv-file-store/storage/file"
	"github.com/sophie-rigg/csv-file-store/worker"
)

var (
	port     int
	workers  int64
	logLevel string
)

func init() {
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.Int64Var(&workers, "workers", 5, "Number of workers")
	flag.StringVar(&logLevel, "log-level", "info", "Log level (debug, info, warn, error, fatal, panic)")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	logger := log.With().Ctx(ctx).Fields(map[string]interface{}{
		"log_level": logLevel,
		"port":      port,
	}).Logger()

	// Set the log level
	l, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		logger.Fatal().Err(err).Msg("parsing log level")
	}

	zerolog.SetGlobalLevel(l)

	// Allows for the storage to be swapped out for gcs, s3, etc
	getStorage := func() storage.Storage {
		return file.New("./files")
	}

	// Get the existing files
	existingFiles, err := getStorage().ListFiles()
	if err != nil {
		logger.Fatal().Err(err).Msg("listing files")
	}

	// Create the cache with the existing files
	localCache, err := cache.NewCache(existingFiles)
	if err != nil {
		logger.Fatal().Err(err).Msg("creating cache")
	}

	processor := worker.New(ctx, workers, localCache, getStorage)
	defer processor.Close()

	// Register the handlers
	router := server.Register(processor, getStorage)

	logger.Info().Msg("starting server")
	// Start the server
	if err = http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		logger.Fatal().Err(err).Msg("server error")
	}
}
