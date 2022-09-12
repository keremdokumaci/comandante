package comandante

import (
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

func Configure(cfg Config) *Comandante {
	var str storage.Storer

	switch cfg.StoreIn {
	case storage.StorageFile:
		str = storage.NewFileStorage()
	case storage.StorageRedis:
		str = storage.NewRedisStorage(cfg.RedisOptions)
	default:
		return nil
	}

	return &Comandante{
		Storage: str,
	}
}
