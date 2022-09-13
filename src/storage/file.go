package storage

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"

	"github.com/keremdokumaci/comandante/src/constants"
	"github.com/keremdokumaci/comandante/src/models"
	"github.com/keremdokumaci/comandante/src/utils"
)

type FileStorage struct{}

func NewFileStorage() *FileStorage {
	return &FileStorage{}
}

func (fileStorage *FileStorage) Write(key string, value string) error {
	err := utils.CreateFileIfNotExists(constants.ConfigurationJsonPath)
	if err != nil {
		return err
	}

	cfgVar := models.ConfigurationVariable{
		Key:           key,
		Value:         value,
		LastUpdatedAt: time.Now(),
	}

	configVariables, err := fileStorage.GetAll()
	if err != nil {
		return err
	}

	if configVariables.GetByKey(key) != nil {
		return ErrConfigurationVariableAlreadyExists
	}

	configVariables = append(configVariables, cfgVar)

	file, err := os.OpenFile(constants.ConfigurationJsonPath, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return errors.New("an error occured while opening configuration json")
	}
	defer file.Close()

	bytes, _ := json.Marshal(configVariables)
	_, err = file.Write(bytes)
	return err
}

func (*FileStorage) GetAll() (models.ArrConfigurationVariable, error) {
	file, err := os.Open(constants.ConfigurationJsonPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var envVars models.ArrConfigurationVariable
	if err = json.Unmarshal(byteValue, &envVars); err != nil {
		return nil, err
	}

	return envVars, nil
}

func (rs *FileStorage) Delete(key string) error {
	return nil //TODO
}

func (rs *FileStorage) Update(key string, newValue string) error {
	return nil //TODO
}
