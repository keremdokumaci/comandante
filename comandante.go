package comandante

import (
	"time"

	"github.com/keremdokumaci/comandante/src/storage"
)

type Config struct {
	SetEnv         EnvVarSetterFunc    `validate:"required"`
	ErrorHandler   ErrorHandler        `validate:"required"`
	RetryTimeInSec int                 `validate:"required"`
	StoreIn        storage.StorageType `validate:"required"`
}

type Comandante struct {
	Storage storage.Storer
}

type EnvVarSetterFunc func(envVars map[string]string)
type ErrorHandler func(err error)

func Configure(cfg Config) *Comandante {
	// TODO: storage exists check
	storage := storage.NewStorage(cfg.StoreIn)
	go func() {
		for {
			definedEnvVars, err := storage.GetAll()
			if err != nil {
				cfg.ErrorHandler(err)
				break
			}

			cfg.SetEnv(definedEnvVars.GetKeyValueMap())
			time.Sleep(time.Duration(cfg.RetryTimeInSec) * time.Second)
		}
	}()

	return &Comandante{
		Storage: storage,
	}
}
