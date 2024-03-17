package csv

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"
)

func TestClient_Write(t *testing.T) {
	file, err := os.Open("../../test-files/test.csv")
	if err != nil {
		t.Errorf("error opening file: %v", err)
	}
	tests := []struct {
		name       string
		Writer     io.ReadWriter
		Reader     io.Reader
		wantResult string
		wantErr    bool
	}{
		{
			name:       "test valid csv",
			Writer:     &bytes.Buffer{},
			Reader:     file,
			wantResult: "name,email,age,email_exists\nsophie,sophie.rigg@gmail.com,25,true\nsophie,sophie.rigg,25,false\n",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := New(context.Background(), tt.Writer, tt.Reader)
			if err != nil {
				t.Errorf("New() error = %v", err)
				return
			}
			if err := c.Write(); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
			c.Close()

			results, err := io.ReadAll(tt.Writer)
			if err != nil {
				t.Errorf("error reading results: %v", err)
			}
			if string(results) == "" {
				t.Errorf("expected non-empty results")
			}
			if string(results) != tt.wantResult {
				t.Errorf("expected \n%s\n, got \n%s", tt.wantResult, string(results))
			}
		})
	}
}
