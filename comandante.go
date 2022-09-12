package comandante

import (
	"github.com/keremdokumaci/comandante/src/storage"
)

type Config struct {
	ErrorHandler ErrorHandler        `validate:"required"`
	StoreIn      storage.StorageType `validate:"required"`
}

type Comandante struct {
	Storage storage.Storer
}

type ErrorHandler func(err error)

func Configure(cfg Config) *Comandante {
	// TODO: storage exists check
	storage := storage.NewStorage(cfg.StoreIn)
	return &Comandante{
		Storage: storage,
	}
}
