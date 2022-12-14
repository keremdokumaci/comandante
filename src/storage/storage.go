package storage

import (
	"errors"

	"github.com/keremdokumaci/comandante/src/models"
)

var (
	ErrConfigurationVariableAlreadyExists = errors.New("key already exists")
)

type Storer interface {
	Get(key string) (*models.ConfigurationVariable, error)
	Write(key string, value string) error
	GetAll() (models.ArrConfigurationVariable, error)
	Delete(key string) error
	Update(key string, newValue string) error
}

type StorageType string

const (
	StorageRedis StorageType = "redis"
)
