package worker

import (
	"testing"
)

func TestJob_IncrementAttempts(t *testing.T) {
	t.Run("increment increases attempts", func(t *testing.T) {
		j := NewJob()
		j.IncrementAttempts()
		if j.GetAttempts() != 1 {
			t.Errorf("expected 1, got %d", j.GetAttempts())
		}
		if j.GetID() == "" {
			t.Errorf("expected non-empty id")
		}
		j.AddData([]byte("data"))
		if string(j.GetData()) != "data" {
			t.Errorf("expected data, got %s", j.GetData())
		}
	})
}
