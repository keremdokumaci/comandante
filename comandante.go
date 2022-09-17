package comandante

import (
	"encoding/json"

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
	case storage.StorageRedis:
		str = storage.NewRedisStorage(cfg.RedisOptions)
	default:
		return nil
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
