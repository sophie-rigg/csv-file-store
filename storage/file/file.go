package file

import (
	"io"
	"os"

	"github.com/sophie-rigg/csv-file-store/storage"
)

var _ storage.Storage = (*client)(nil)

type client struct {
	io.ReadWriteCloser
	directory string
}

func New(directory string) *client {
	return &client{
		directory: directory,
	}
}

func (c *client) CreateFile(id string) error {
	file, err := os.Create(c.getFileName(id))
	if err != nil {
		return err
	}
	c.ReadWriteCloser = file
	return nil
}

func (c *client) OpenFile(id string) error {
	file, err := os.Open(c.getFileName(id))
	if err != nil {
		return err
	}
	c.ReadWriteCloser = file
	return nil
}

func (c *client) RemoveFile(id string) error {
	return os.Remove(c.getFileName(id))
}

// ListFiles returns a list of all files in the directory for initialization of the cache
func (c *client) ListFiles() ([]string, error) {
	dir, err := os.Open(c.directory)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	files, err := dir.Readdirnames(0)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// getFileName returns the full path to the file
func (c *client) getFileName(id string) string {
	return c.directory + "/" + id + ".csv"
}
