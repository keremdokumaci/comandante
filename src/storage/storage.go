package storage

import (
	"github.com/keremdokumaci/comandante/src/models"
)

type Storer interface {
	Write(key string, value string) error
	GetAll() (models.ArrConfigurationVariable, error)
}

type StorageType string

const (
	StorageFile  StorageType = "file"
	StorageRedis StorageType = "redis"
)
