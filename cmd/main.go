package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/keremdokumaci/comandante"
	"github.com/keremdokumaci/comandante/src/storage"
)

func main() {
	cmdt := comandante.Configure(comandante.Config{
		ErrorHandler: func(err error) {
			fmt.Println(err.Error())
		},
		StoreIn: storage.StorageRedis,
		RedisOptions: &redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	})

	http.HandleFunc("/comandante", cmdt.HandlerFunc)

	http.HandleFunc("/config_variables", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(os.Getenv("comandante"))) //nolint
	})

	http.ListenAndServe(":8080", nil) //nolint
}
