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

var (
	ErrConfigurationVariableAlreadyExists = errors.New("key already exists")
	ErrKeyAndValueFieldsAreRequired       = errors.New("key and value fields are required")
)

type FileStorage struct{}

func NewFileStorage() *FileStorage {
	return &FileStorage{}
}

func (fileStorage *FileStorage) Write(key string, value string) error {
	if key == "" || value == "" {
		return ErrKeyAndValueFieldsAreRequired
	}

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
	json.Unmarshal(byteValue, &envVars)

	return envVars, nil
}