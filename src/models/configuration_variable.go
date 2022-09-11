package models

import "time"

type ConfigurationVariable struct {
	Key           string
	Value         string
	LastUpdatedAt time.Time
}

type ArrConfigurationVariable []ConfigurationVariable

func (arr ArrConfigurationVariable) GetKeyValueMap() map[string]string {
	var keyValueMap = make(map[string]string)
	for _, el := range arr {
		keyValueMap[el.Key] = el.Value
	}

	return keyValueMap
}

func (arr ArrConfigurationVariable) GetByKey(key string) *ConfigurationVariable {
	for _, el := range arr {
		if el.Key == key {
			return &el
		}
	}
	return nil
}
