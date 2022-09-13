package client

import (
	"html/template"

	"github.com/keremdokumaci/comandante/src/models"
)

type PageData struct {
	ConfigVariables models.ArrConfigurationVariable
	AddNewConfig    func()
}

func readHtml() (string, error) {
	return htmlContent, nil
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
