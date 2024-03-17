package server

import (
	"github.com/gorilla/mux"
	"github.com/sophie-rigg/csv-file-store/server/download"
	"github.com/sophie-rigg/csv-file-store/server/upload"
	"github.com/sophie-rigg/csv-file-store/storage"
	"github.com/sophie-rigg/csv-file-store/worker"
)

// Register registers the API routes for the server
func Register(processor worker.Client, newStorage func() storage.Storage) *mux.Router {
	router := mux.NewRouter()

	router.Path("/API/upload").Handler(upload.NewHandler(processor))
	// id is the file id
	router.Path("/API/download/{id}").Handler(download.NewHandler(newStorage, processor.GetCache()))

	return router
}
