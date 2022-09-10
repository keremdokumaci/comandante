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
			keyValueMap := make(map[string]string)
			for key, value := range definedEnvVars {
				keyValueMap[key] = value.Value
			}
			cfg.SetEnv(keyValueMap)
			time.Sleep(3 * time.Second)
		}
	}()
}
