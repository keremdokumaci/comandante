package comandante

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/keremdokumaci/comandante/src/client"
)

func (c *Comandante) renderPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/plain")
	envVars, _ := c.Storage.GetAll() //TODO: log error here

	htmlData := client.PageData{
		ConfigVariables: envVars,
	}

	template, err := client.GenerateTemplate(htmlData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = template.Execute(w, htmlData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (c *Comandante) addConfig(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request client.AddConfigurationVariableRequestModel
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

	os.Setenv(request.Key, request.Value)

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
