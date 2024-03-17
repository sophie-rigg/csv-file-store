package upload

import (
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sophie-rigg/csv-file-store/worker"
)

type handler struct {
	processor worker.Client
	logger    zerolog.Logger
}

// handler must implement the http.Handler interface
var _ http.Handler = (*handler)(nil)

func NewHandler(processor worker.Client) *handler {
	return &handler{
		processor: processor,
		logger: log.With().Fields(map[string]interface{}{
			"handler":      "upload",
			"storage_type": "file",
		}).Logger(),
	}
}

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		h.handlePost(writer, request)
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *handler) handlePost(writer http.ResponseWriter, request *http.Request) {
	// Create a new job with a new id
	job := worker.NewJob()
	// Add the job to the cache as uncompleted
	h.processor.GetCache().Add(job.GetID())

	data, err := io.ReadAll(request.Body)
	if err != nil {
		h.logger.Error().Err(err).Msg("error reading request body")
		h.processor.GetCache().Remove(job.GetID())
		http.Error(writer, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Add the data to the job
	job.AddData(data)
	// Queue the job for processing
	h.processor.AddJob(job)

	response, err := newUploadPostResponse(job.GetID()).MarshalToJson()
	if err != nil {
		http.Error(writer, fmt.Sprintf("Error creating response, id: %s", job.GetID()), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(response)
	if err != nil {
		h.logger.Error().Err(err).Msg("error writing response")
	}
}
