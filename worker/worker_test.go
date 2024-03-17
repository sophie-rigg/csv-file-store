package worker

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sophie-rigg/csv-file-store/cache"
	"github.com/sophie-rigg/csv-file-store/storage"
	mock_storage "github.com/sophie-rigg/csv-file-store/storage/mocks"
)

func TestWorker_processJob(t *testing.T) {
	type fields struct {
		workers int64
		storage func(client *mock_storage.MockStorage)
	}
	tests := []struct {
		name    string
		fields  fields
		jobs    []*Job
		wantErr bool
	}{
		{
			name: "process job",
			fields: fields{
				workers: 5,
				storage: func(client *mock_storage.MockStorage) {
					client.EXPECT().CreateFile("123").Return(nil)
					client.EXPECT().Write([]byte("name,email,age,email_exists\nsophie,sophie.rigg@gmail.com,25,true\nsophie,sophie.rigg,25,false\n")).Return(len([]byte("test")), nil)
					client.EXPECT().Close().Return(nil)
				},
			},
			jobs: []*Job{
				{
					id:   "123",
					data: []byte("name,email,age\nsophie,sophie.rigg@gmail.com,25\nsophie,sophie.rigg,25"),
				},
			},
		},
		{
			name: "process job fails once but succeeds on retry",
			fields: fields{
				workers: 5,
				storage: func(client *mock_storage.MockStorage) {
					client.EXPECT().CreateFile("123").Return(errors.New("test error"))
					client.EXPECT().CreateFile("123").Return(nil)
					client.EXPECT().Write([]byte("name,email,age,email_exists\nsophie,sophie.rigg@gmail.com,25,true\nsophie,sophie.rigg,25,false\n")).Return(len([]byte("test")), nil)
					client.EXPECT().Close().Return(nil)
				},
			},
			jobs: []*Job{
				{
					id:   "123",
					data: []byte("name,email,age\nsophie,sophie.rigg@gmail.com,25\nsophie,sophie.rigg,25"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := cache.NewCache([]string{})
			if err != nil {
				t.Errorf("error creating cache: %v", err)
			}
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			s := mock_storage.NewMockStorage(ctrl)
			tt.fields.storage(s)
			newStorage := func() storage.Storage {
				return s
			}

			w := New(context.Background(), tt.fields.workers, c, newStorage)

			for _, job := range tt.jobs {
				w.AddJob(job)
			}

			var wg sync.WaitGroup
			for _, j := range tt.jobs {
				wg.Add(1)
				go func(job *Job) {
					for {
						complete, _ := w.GetCache().IsJobCompleted(job.GetID())
						if complete {
							wg.Done()
							break
						}
					}
				}(j)
			}
			wg.Wait()
		})
	}
}
