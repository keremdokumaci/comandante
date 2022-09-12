package storage

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/keremdokumaci/comandante/src/models"
)

const keyPrefix = "comandante"

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(options *redis.Options) *RedisStorage {
	return &RedisStorage{
		client: redis.NewClient(options),
	}
}

func (rs *RedisStorage) Write(key string, value string) error {
	ctx := context.Background()

	cfgVar := models.ConfigurationVariable{
		Key:           key,
		Value:         value,
		LastUpdatedAt: time.Now(),
	}

	marshaledCfgVar, err := json.Marshal(cfgVar)
	if err != nil {
		return err
	}

	err = rs.client.Set(ctx, getKeyWithPrefix(key), marshaledCfgVar, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rs *RedisStorage) GetAll() (models.ArrConfigurationVariable, error) {
	var configVars models.ArrConfigurationVariable

	ctx := context.Background()
	iter := rs.client.Scan(ctx, 0, keyPrefix+":*", 0).Iterator()
	for iter.Next(ctx) {
		val, err := rs.client.Get(ctx, iter.Val()).Result()
		if err != nil {
			return nil, err
		}

		var configVariable models.ConfigurationVariable
		err = json.Unmarshal([]byte(val), &configVariable)
		if err != nil {
			return nil, err
		}

		configVars = append(configVars, configVariable)
	}

	return configVars, nil
}

func getKeyWithPrefix(key string) string {
	return keyPrefix + ":" + key
}
