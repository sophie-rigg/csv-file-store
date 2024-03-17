package upload

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sophie-rigg/csv-file-store/cache"
	"github.com/sophie-rigg/csv-file-store/worker/mocks"
)

func Test_handler_ServeHTTP_Success(t *testing.T) {
	testFile, err := os.Open("../../test-files/test.csv")
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest("POST", "http://localhost:8080", testFile)

	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	workerClient := mocks.NewMockClient(ctrl)

	localCache, err := cache.NewCache([]string{})
	if err != nil {
		t.Fatal(err)
	}
	workerClient.EXPECT().GetCache().Return(localCache)
	workerClient.EXPECT().AddJob(gomock.Any())

	NewHandler(workerClient).ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Fatal("expected 200 status code, got: ", resp.StatusCode)
	}
}
