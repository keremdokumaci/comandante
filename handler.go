package comandante

import (
	"html/template"
	"net/http"

	filemanager "github.com/keremdokumaci/comandante/src/file_manager"
)

type envVar struct {
	Key   string
	Value string
}
type htmlData struct {
	EnvVars []envVar
}

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	htmlData := htmlData{}
	envVars := filemanager.ReadConfigurationJson()
	for key, value := range envVars {
		envVar := envVar{
			Key:   key,
			Value: value,
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
