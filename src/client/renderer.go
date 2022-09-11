package client

import (
	"io"
	"os"

	"github.com/keremdokumaci/comandante/src/constants"
)

func ReadHtml() (string, error) {
	file, err := os.Open(constants.ComandanteHtmlPath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(byteValue), nil
}
