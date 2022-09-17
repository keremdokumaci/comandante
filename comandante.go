package comandante

import (
	"encoding/json"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/keremdokumaci/comandante/src/storage"
)

type Config struct {
	ErrorHandler ErrorHandler        `validate:"required"`
	StoreIn      storage.StorageType `validate:"required"`
	RedisOptions *redis.Options
}

type Comandante struct {
	Storage storage.Storer
}

type ErrorHandler func(err error)

var str storage.Storer

func Configure(cfg Config) *Comandante {

	switch cfg.StoreIn {
	case storage.StorageFile:
		str = storage.NewFileStorage()
	case storage.StorageRedis:
		str = storage.NewRedisStorage(cfg.RedisOptions)
	default:
		return nil
	}

	vars, _ := str.GetAll()
	for key, value := range vars.GetKeyValueMap() {
		os.Setenv(key, value)
	}

	return &Comandante{
		Storage: str,
	}
}

func Get[T any](key string) (*T, error) {
	val, err := str.Get(key)
	if err != nil {
		return nil, err
	}

	var genericType T
	err = json.Unmarshal([]byte(val.Value), &genericType)
	if err != nil {
		return nil, err
	}

	return &genericType, nil
}
