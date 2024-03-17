package file

import (
	"os"
	"testing"
)

func Test_client_CreateFile(t *testing.T) {
	id := "123"
	err := os.Mkdir("./temp-files", 0755)
	if err != nil {
		t.Errorf("Creating directory error = %v", err)
	}
	c := New("./temp-files")
	err = c.CreateFile(id)
	if err != nil {
		t.Errorf("CreateFile() error = %v", err)
	}
	err = c.OpenFile(id)
	if err != nil {
		t.Errorf("OpenFile() error = %v", err)
	}
	files, err := c.ListFiles()
	if err != nil {
		t.Errorf("ListFiles() error = %v", err)
	}
	if len(files) != 1 {
		t.Errorf("ListFiles() error = %v", files)
	}

	err = c.RemoveFile(id)
	if err != nil {
		t.Errorf("RemoveFile() error = %v", err)
	}

	err = os.RemoveAll("./temp-files")
	if err != nil {
		t.Errorf("Removing directory error = %v", err)
	}
}
