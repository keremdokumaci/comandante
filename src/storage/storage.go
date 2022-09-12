package storage

import (
	"errors"

	"github.com/keremdokumaci/comandante/src/models"
)

var (
	ErrConfigurationVariableAlreadyExists = errors.New("key already exists")
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
