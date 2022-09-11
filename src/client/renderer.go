package client

import (
	"html/template"
	"io"
	"os"

	"github.com/keremdokumaci/comandante/src/constants"
	"github.com/keremdokumaci/comandante/src/models"
)

type PageData struct {
	ConfigVariables models.ArrConfigurationVariable
	AddNewConfig    func()
}

func readHtml() (string, error) {
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

func GenerateTemplate(data PageData) (*template.Template, error) {
	htmlContent, err := readHtml()
	if err != nil {
		return nil, err
	}

	t, err := template.New("comandante").Parse(htmlContent)
	if err != nil {
		return nil, err
	}

	return t, nil
}
