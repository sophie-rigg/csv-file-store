package download

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sophie-rigg/csv-file-store/cache"
	"github.com/sophie-rigg/csv-file-store/storage"
	mock_storage "github.com/sophie-rigg/csv-file-store/storage/mocks"
)

func Test_handler_ServeHTTP_Success(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/API/download/123", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "a4966f92-890e-4d0f-8e79-769a556ae57d"})

	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	s := mock_storage.NewMockStorage(ctrl)
	s.EXPECT().OpenFile("a4966f92-890e-4d0f-8e79-769a556ae57d").Return(nil)
	s.EXPECT().Read(gomock.Any()).DoAndReturn(func(p []byte) (int, error) {
		copy(p, "file content")
		return 11, nil
	})
	s.EXPECT().Read(gomock.Any()).Return(0, io.EOF)

	localCache, err := cache.NewCache([]string{"a4966f92-890e-4d0f-8e79-769a556ae57d.csv"})
	if err != nil {
		t.Fatal(err)
	}

	NewHandler(func() storage.Storage {
		return s
	}, localCache).ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Fatal("expected 200 status code, got: ", resp.StatusCode)
	}
}

func Test_handler_ServeHTTP_Job_Incomplete(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/API/download/123", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "a4966f92-890e-4d0f-8e79-769a556ae57d"})

	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	s := mock_storage.NewMockStorage(ctrl)

	localCache, err := cache.NewCache([]string{})
	if err != nil {
		t.Fatal(err)
	}

	localCache.Add("a4966f92-890e-4d0f-8e79-769a556ae57d")

	NewHandler(func() storage.Storage {
		return s
	}, localCache).ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != 423 {
		t.Fatal("expected 423 status code, got: ", resp.StatusCode)
	}
}

func Test_handler_ServeHTTP_No_ID(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/API/download", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	s := mock_storage.NewMockStorage(ctrl)

	localCache, err := cache.NewCache([]string{"123.csv"})
	if err != nil {
		t.Fatal(err)
	}

	NewHandler(func() storage.Storage {
		return s
	}, localCache).ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != 400 {
		t.Fatal("expected 400 status code, got: ", resp.StatusCode)
	}
}

func Test_handler_ServeHTTP_Bad_ID(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/API/download", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "123"})

	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	s := mock_storage.NewMockStorage(ctrl)

	localCache, err := cache.NewCache([]string{})
	if err != nil {
		t.Fatal(err)
	}

	NewHandler(func() storage.Storage {
		return s
	}, localCache).ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != 400 {
		t.Fatal("expected 400 status code, got: ", resp.StatusCode)
	}
}

func Test_handler_ServeHTTP_Bad_Method(t *testing.T) {
	req := httptest.NewRequest("PUT", "http://localhost:8080/API/download", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	s := mock_storage.NewMockStorage(ctrl)

	localCache, err := cache.NewCache([]string{})
	if err != nil {
		t.Fatal(err)
	}

	NewHandler(func() storage.Storage {
		return s
	}, localCache).ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != 405 {
		t.Fatal("expected 405 status code, got: ", resp.StatusCode)
	}
}
