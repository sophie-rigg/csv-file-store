package worker

import (
	"bytes"
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sophie-rigg/csv-file-store/cache"
	"github.com/sophie-rigg/csv-file-store/storage"
	"github.com/sophie-rigg/csv-file-store/writer/csv"
	"golang.org/x/sync/semaphore"
)

//go:generate mockgen -destination=mocks/mock_worker.go -package=mocks --source=worker.go
type Client interface {
	AddJob(job *Job)
	GetCache() *cache.Cache
}

type Worker struct {
	ctx        context.Context
	jobs       chan *Job
	sem        *semaphore.Weighted
	localCache *cache.Cache
	newStorage func() storage.Storage
	logger     zerolog.Logger
}

func New(ctx context.Context, workers int64, localCache *cache.Cache, newStorage func() storage.Storage) *Worker {
	w := &Worker{
		ctx:        ctx,
		jobs:       make(chan *Job, 20),
		sem:        semaphore.NewWeighted(workers),
		localCache: localCache,
		newStorage: newStorage,
		logger: log.With().Ctx(ctx).Fields(map[string]interface{}{
			"component": "worker",
			"workers":   workers,
		}).Logger(),
	}
	// Start the processor
	go w.Start()

	return w
}

func (w *Worker) Start() {
	for job := range w.jobs {
		ctx, cancel := context.WithTimeout(w.ctx, 2*time.Minute)

		// Fail the job if it has been attempted 3 times
		if job.GetAttempts() >= 3 {
			w.logger.Error().Str("id", job.GetID()).Msg("job failed after 3 attempts")
			w.failJob(job)
			cancel()
			continue
		}

		// Acquire a semaphore to limit the number of concurrent jobs
		err := w.sem.Acquire(w.ctx, 1)
		if err != nil {
			w.logger.Error().Err(err).Msg("error acquiring semaphore")
			w.retryJob(job)
			cancel()
			continue
		}
		go func(j *Job) {
			defer func() {
				w.sem.Release(1)
				cancel()
			}()
			err = w.processJob(ctx, j)
			if err != nil {
				w.logger.Error().Err(err).Str("id", j.GetID()).Msg("error processing job")
				w.retryJob(j)
			}
		}(job)
	}
}

// failJob removes the job from the cache and storage
func (w *Worker) failJob(job *Job) {
	w.localCache.Remove(job.GetID())
	// remove file from storage if it exists ignore error as it may not exist
	w.newStorage().RemoveFile(job.GetID())
}

func (w *Worker) retryJob(job *Job) {
	job.IncrementAttempts()
	w.AddJob(job)
}

// AddJob adds a job to the worker, called by other components
func (w *Worker) AddJob(job *Job) {
	w.jobs <- job
}

func (w *Worker) Close() {
	close(w.jobs)
}

func (w *Worker) processJob(ctx context.Context, job *Job) error {
	store := w.newStorage()

	err := store.CreateFile(job.GetID())
	if err != nil {
		return err
	}

	csvWriter, err := csv.New(ctx, store, bytes.NewReader(job.GetData()))
	if err != nil {
		return err
	}

	err = csvWriter.Write()
	if err != nil {
		return err
	}

	csvWriter.Close()

	err = store.Close()
	if err != nil {
		return err
	}
	// Mark the job as completed
	w.localCache.JobCompleted(job.GetID())
	return nil
}

func (w *Worker) GetCache() *cache.Cache {
	return w.localCache
}
