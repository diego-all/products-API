package main

import (
	"crypto/tls"
	"fmt"

	models "github.com/diego-all/products-API/internal"

	"github.com/diego-all/products-API/database"

	"log"
	"net/http"
	"os"
)

type config struct {
	port     int
	certPath string
	keyPath  string
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	models   models.Models
}

const (
	DSN = "data.sqlite"
)

func main() {

	var cfg config
	cfg.port = 9090

	cfg.certPath = "/home/diegoall/MAESTRIA_ING/domain-model/products-API/cmd/api/server.pem"
	cfg.keyPath = "/home/diegoall/MAESTRIA_ING/domain-model/products-API/cmd/api/server.key"

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := database.ConnectSQLite(DSN)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	defer db.SQL.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		models:   models.New(db.SQL),
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

func (app *application) serve() error {
	app.infoLog.Println("API listening on port", app.config.port)

	// Configuramos el servidor TLS para solo aceptar TLS 1.3
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13, // Versión mínima TLS 1.3
		MaxVersion: tls.VersionTLS13, // Versión máxima también TLS 1.3
	}

	srv := &http.Server{
		Addr:      fmt.Sprintf(":%d", app.config.port),
		Handler:   app.routes(),
		TLSConfig: tlsConfig,
	}
	//return srv.ListenAndServe()
	return srv.ListenAndServeTLS(app.config.certPath, app.config.keyPath)
}
