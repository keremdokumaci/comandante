package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/keremdokumaci/comandante"
	"github.com/keremdokumaci/comandante/src/storage"
)

func main() {
	cmdt := comandante.Configure(comandante.Config{
		ErrorHandler: func(err error) {
			fmt.Println(err.Error())
		},
		RetryTimeInSec: 3,
		StoreIn:        storage.StorageFile,
	})

	http.HandleFunc("/comandante", cmdt.HandlerFunc)

	http.HandleFunc("/config_variables", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(os.Getenv("comandante")))
	})

	http.ListenAndServe(":8080", nil) //nolint
}
