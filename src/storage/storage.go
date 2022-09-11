package storage

import "github.com/keremdokumaci/comandante/src/models"

type Storer interface {
	Write(key string, value string) error
	GetAll() (models.ArrConfigurationVariable, error)
}

type StorageType string

const (
	StorageFile  StorageType = "file"
	StorageRedis StorageType = "redis"
)

func NewStorage(storageType StorageType) Storer {
	switch storageType {
	case StorageFile:
		return NewFileStorage()
	case StorageRedis:
		return nil
	default:
		return nil
	}
}
