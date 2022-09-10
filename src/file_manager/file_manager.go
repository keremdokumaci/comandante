package filemanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

const (
	configurationJsonPath = "./server_configuration.json"
	comandanteHtmlPath    = "./public/comandante.html"
)

var (
	ErrKeyAlreadyExists = errors.New("key already exists")
)

type ConfigVar struct {
	Value         string
	LastUpdatedAt time.Time
}

func CreateFileIfNotExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func Write(key string, value string) error {
	if key == "" || value == "" {
		return errors.New("invalid request parameters")
	}

	err := CreateFileIfNotExists(configurationJsonPath)
	if err != nil {
		return err
	}

	configVars := ReadConfigurationJson()
	if configVars[key].Value != "" {
		return ErrKeyAlreadyExists
	}

	configVars[key] = ConfigVar{
		Value:         value,
		LastUpdatedAt: time.Now(),
	}

	file, err := os.OpenFile(configurationJsonPath, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return errors.New("an error occured while opening configuration json")
	}
	defer file.Close()

	bytes, _ := json.Marshal(configVars)
	_, err = file.Write(bytes)
	return err
}

func ReadConfigurationJson() map[string]ConfigVar {
	file, err := os.Open(configurationJsonPath)
	if err != nil {
		fmt.Println("An error occured while opening configuration json : ", err)
		return map[string]ConfigVar{}
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error : ", err)
		return map[string]ConfigVar{}
	}

	envVars := make(map[string]ConfigVar)
	json.Unmarshal(byteValue, &envVars)

	return envVars
}

func ReadHtml() string {
	file, err := os.Open(comandanteHtmlPath)
	if err != nil {
		fmt.Println("An error occured while opening html file : ", err)
		return ""
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error : ", err)
	}

	return string(byteValue)
}
