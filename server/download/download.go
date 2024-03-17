package download

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sophie-rigg/csv-file-store/cache"
	"github.com/sophie-rigg/csv-file-store/storage"
	"github.com/sophie-rigg/csv-file-store/utils"
)

type handler struct {
	newStorage func() storage.Storage
	localCache *cache.Cache
	logger     zerolog.Logger
}

var (
	_ http.Handler = (*handler)(nil)

	_errNoIDProvided = errors.New("no id provided in query")
)

func NewHandler(newStorage func() storage.Storage, localCache *cache.Cache) *handler {
	return &handler{
		newStorage: newStorage,
		localCache: localCache,
		logger: log.With().Fields(map[string]interface{}{
			"handler":      "download",
			"storage_type": "file",
		}).Logger(),
	}
}

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Only allow GET requests
	switch request.Method {
	case http.MethodGet:
		h.handleGet(writer, request)
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *handler) handleGet(writer http.ResponseWriter, request *http.Request) {
	store := h.newStorage()

	id, err := getIDFromRequest(request)
	if err != nil {
		h.logger.Error().Err(err).Msg("error getting id from request")
		http.Error(writer, "Error getting id from request", http.StatusBadRequest)
		return
	}

	// Check if the id is in the correct format
	if !utils.CheckIDFormat(id) {
		h.logger.Error().Err(err).Msg("invalid id")
		http.Error(writer, fmt.Sprintf("Invalid id: %s", id), http.StatusBadRequest)
		return
	}

	// Check if the job is completed
	complete, err := h.localCache.IsJobCompleted(id)
	if err != nil {
		h.logger.Error().Err(err).Msg("error checking if job is completed")
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if !complete {
		http.Error(writer, "Job not completed", http.StatusLocked)
		return
	}

	// Open the file
	err = store.OpenFile(id)
	if err != nil {
		h.logger.Error().Err(err).Msg("error opening file")
		http.Error(writer, fmt.Sprintf("Error opening file: %s.csv", id), http.StatusBadRequest)
		return
	}

	// Copy the file to the response
	byteCount, err := io.Copy(writer, store)
	if err != nil {
		h.logger.Error().Err(err).Msg("error copying file to response")
		http.Error(writer, "Error copying file to response", http.StatusLocked)
		return
	}

	if byteCount == 0 {
		h.logger.Error().Err(err).Msg("error copying file to response")
		http.Error(writer, fmt.Sprintf("Empty file still uploading: %s.csv", id), http.StatusLocked)
		return
	}
}

// getIDFromRequest gets the id from the request vars
func getIDFromRequest(r *http.Request) (string, error) {
	queryVars := mux.Vars(r)
	id, ok := queryVars["id"]
	if !ok {
		return "", _errNoIDProvided
	}
	return id, nil
}
