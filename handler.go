package comandante

import (
	"net/http"

	filemanager "github.com/keremdokumaci/comandante/src/file_manager"
)

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/plain")
	w.Write([]byte(filemanager.ReadHtml()))
}
