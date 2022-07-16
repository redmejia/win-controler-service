package main

import (
	"log"
	"net/http"
	"os"
	"win/controler/cmd/api"
	"win/controler/cmd/router"
)

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	controler := api.ApiConfig{
		Port:     ":8080",
		Infolog:  infoLog,
		Errorlog: errorLog,
	}

	srv := &http.Server{
		Addr:    controler.Port,
		Handler: router.Rouer(&controler),
	}

	infoLog.Printf("Service running at http://localhost:%s", controler.Port)
	errorLog.Fatal(srv.ListenAndServe())
}
