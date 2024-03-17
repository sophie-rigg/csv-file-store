package cache

import (
	"strings"
	"testing"
)

func TestCache_Add(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		startingKeys []string
	}{
		{
			name: "add key to cache",
			key:  "test",
			startingKeys: []string{
				"a8901259-d453-4ecf-9258-7230e43b1f29.csv",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewCache(tt.startingKeys)
			if err != nil {
				t.Errorf("error creating cache: %v", err)
			}
			for _, key := range tt.startingKeys {
				complete, err := c.IsJobCompleted(strings.Split(key, ".")[0])
				if err != nil {
					t.Errorf("error checking if job is completed: %v", err)
				}
				if !complete {
					t.Errorf("expected job to be completed")
				}
			}
			c.Add(tt.key)
			complete, err := c.IsJobCompleted(tt.key)
			if err != nil {
				t.Errorf("error checking if job is completed: %v", err)
			}
			if complete {
				t.Errorf("expected job to not be completed")
			}
			c.JobCompleted(tt.key)
			complete, err = c.IsJobCompleted(tt.key)
			if err != nil {
				t.Errorf("error checking if job is completed: %v", err)
			}
			if !complete {
				t.Errorf("expected job to be completed")
			}
			c.Remove(tt.key)
			_, err = c.IsJobCompleted(tt.key)
			if err == nil {
				t.Errorf("expected error checking if job is completed")
			}
		})
	}
}
