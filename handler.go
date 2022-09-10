package comandante

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"time"

	filemanager "github.com/keremdokumaci/comandante/src/file_manager"
)

type envVar struct {
	Key           string
	Value         string
	LastUpdatedAt string
}
type htmlData struct {
	EnvVars      []envVar
	AddNewConfig func()
}

type addConfigRequest struct {
	Key   string
	Value string
}

func renderPage(w http.ResponseWriter, r *http.Request) {
	htmlData := htmlData{}
	envVars := filemanager.ReadConfigurationJson()
	for key, value := range envVars {
		envVar := envVar{
			Key:           key,
			Value:         value.Value,
			LastUpdatedAt: value.LastUpdatedAt.Format(time.RFC3339), // TODO: format by timezone
		}
		htmlData.EnvVars = append(htmlData.EnvVars, envVar)
	}

	w.Header().Add("Content Type", "text/plain")
	t, err := template.New("comandante").Parse(filemanager.ReadHtml())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	t.Execute(w, htmlData)
}

func addConfig(w http.ResponseWriter, r *http.Request) {
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

	err = filemanager.Write(request.Key, request.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		renderPage(w, r)
	case http.MethodPost:
		addConfig(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
