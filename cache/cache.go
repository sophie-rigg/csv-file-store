package cache

import (
	"fmt"
	"strings"
	"sync"

	"github.com/sophie-rigg/csv-file-store/utils"
)

type Cache struct {
	sync.Mutex
	data map[string]bool
}

// NewCache creates a new cache with the given files in format uuid.csv
func NewCache(files []string) (*Cache, error) {
	cache := make(map[string]bool)
	for _, file := range files {
		splitFileName := strings.Split(file, ".")
		// If the file is not in the format uuid.csv, skip it
		if len(splitFileName) != 2 || splitFileName[1] != "csv" || !utils.CheckIDFormat(splitFileName[0]) {
			continue
		}
		// Add the uuid to the cache, assuming it was completed
		cache[splitFileName[0]] = true
	}

	return &Cache{
		data: cache,
	}, nil
}

// Add adds a job to the cache uncompleted
func (c *Cache) Add(key string) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = false
}

// JobCompleted marks the job as completed
func (c *Cache) JobCompleted(key string) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = true
}

// IsJobCompleted returns true if the job is completed, false if it is not, and an error if the job is not found
func (c *Cache) IsJobCompleted(key string) (bool, error) {
	c.Lock()
	defer c.Unlock()
	complete, ok := c.data[key]
	if !ok {
		return false, fmt.Errorf("file %s.csv not found", key)
	}
	return complete, nil
}

// Remove removes a job from the cache
func (c *Cache) Remove(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.data, key)
}
