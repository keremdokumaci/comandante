package main

import (
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
	})
	http.HandleFunc("/comandante", comandante.HandlerFunc)

	// http.HandleFunc("/variable", func(w http.ResponseWriter, r *http.Request) {
	// 	envVar := os.Getenv("example_env_var")
	// 	w.Write([]byte(envVar))
	// })

	http.ListenAndServe(":8080", nil)
}
