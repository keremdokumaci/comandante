package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/keremdokumaci/comandante"
)

func customSetterFunc(envVars map[string]string) {
	for key, value := range envVars {
		os.Setenv(key, value)
	}
}

func main() {
	comandante.Configure(comandante.Config{
		SetEnv: customSetterFunc,
		ErrorHandler: func(err error) {
			fmt.Println(err.Error())
		},
		RetryTimeInSec: 3,
	})

	http.HandleFunc("/comandante", comandante.HandlerFunc)
	http.ListenAndServe(":8080", nil)
}
