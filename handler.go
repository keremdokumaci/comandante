package comandante

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"

	"github.com/keremdokumaci/comandante/src/client"
	"github.com/keremdokumaci/comandante/src/models"
)

type htmlData struct {
	ConfigVariables models.ArrConfigurationVariable
	AddNewConfig    func()
}

type addConfigRequest struct {
	Key   string
	Value string
}

func (c *Comandante) renderPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/plain")
	htmlData := htmlData{}

	envVars, _ := c.Storage.GetAll() //TODO: log error here

	htmlData.ConfigVariables = envVars

	htmlContent, err := client.ReadHtml()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t, err := template.New("comandante").Parse(htmlContent)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, htmlData)
}

func (c *Comandante) addConfig(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request addConfigRequest
	err = json.Unmarshal(bytes, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.Storage.Write(request.Key, request.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *Comandante) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.renderPage(w, r)
	case http.MethodPost:
		c.addConfig(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
