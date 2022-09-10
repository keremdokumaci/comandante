package comandante

import (
	"time"

	filemanager "github.com/keremdokumaci/comandante/src/file_manager"
)

type Config struct {
	SetEnv         EnvVarSetterFunc `validate:"required"`
	ErrorHandler   ErrorHandler     `validate:"required"`
	RetryTimeInSec int              `validate:"required"`
}

type EnvVarSetterFunc func(envVars map[string]string)
type ErrorHandler func(err error)

func Configure(cfg Config) {
	go func() {
		for {
			definedEnvVars, err := filemanager.ReadConfigurationJson()
			if err != nil {
				cfg.ErrorHandler(err)
				break
			}

			keyValueMap := make(map[string]string)
			for key, value := range definedEnvVars {
				keyValueMap[key] = value.Value
			}
			cfg.SetEnv(keyValueMap)
			time.Sleep(time.Duration(cfg.RetryTimeInSec) * time.Second)
		}
	}()
}
