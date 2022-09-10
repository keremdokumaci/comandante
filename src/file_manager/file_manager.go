package filemanager

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	configurationJsonPath = "./server_configuration.json"
	comandanteHtmlPath    = "./public/comandante.html"
)

func Write() error {
	fmt.Println(configurationJsonPath)
	return nil
}

func ReadConfigurationJson() map[string]string {
	file, err := os.Open(configurationJsonPath)
	if err != nil {
		fmt.Println("An error occured while opening configuration json : ", err)
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error : ", err)
	}

	envVars := make(map[string]string)
	json.Unmarshal(byteValue, &envVars)

	return envVars
}

func ReadHtml() string {
	file, err := os.Open(comandanteHtmlPath)
	if err != nil {
		fmt.Println("An error occured while opening html file : ", err)
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error : ", err)
	}

	return string(byteValue)
}
