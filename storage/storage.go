package storage

import (
	"io"
)

// Storage is an interface for a storage system
//
//go:generate mockgen -destination=mocks/mock_storage.go --source=storage.go package=storage Storage
type Storage interface {
	io.ReadWriteCloser
	CreateFile(id string) error
	OpenFile(id string) error
	RemoveFile(id string) error
	ListFiles() ([]string, error)
}
