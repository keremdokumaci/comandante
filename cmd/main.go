package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/keremdokumaci/comandante"
	"github.com/keremdokumaci/comandante/src/storage"
)

func customSetterFunc(envVars map[string]string) {
	for key, value := range envVars {
		os.Setenv(key, value)
	}
}

func main() {
	cmdt := comandante.Configure(comandante.Config{
		SetEnv: customSetterFunc,
		ErrorHandler: func(err error) {
			fmt.Println(err.Error())
		},
		RetryTimeInSec: 3,
		StoreIn:        storage.StorageFile,
	})

	http.HandleFunc("/comandante", cmdt.HandlerFunc)
	http.ListenAndServe(":8080", nil)
}
