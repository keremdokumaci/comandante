package comandante

import (
	"time"

	filemanager "github.com/keremdokumaci/comandante/src/file_manager"
)

type Config struct {
	SetEnv EnvVarSetterFunc `validate:"required"`
}

type EnvVarSetterFunc func(envVars map[string]string)

func Configure(cfg Config) {
	go func() {
		for {
			definedEnvVars := filemanager.ReadConfigurationJson()
			cfg.SetEnv(definedEnvVars)
			time.Sleep(3 * time.Second)
		}
	}()
}
