package main

import (
	"TravelService/src/config"
	"TravelService/src/database"
	"TravelService/src/service"
	"TravelService/src/service/auth"
	"flag"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
)

func main() {
	configFile := flag.String("config", "src/config.json", "Configuration file in JSON-format")
	flag.Parse()

	if len(*configFile) > 0 {
		config.FilePath = *configFile
	}

	err := config.Load()
	if err != nil {
		log.Fatalf("error while reading config: %s", err)
	}

	f, err := os.OpenFile(config.Config.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	database.DB, err = database.SetupPostgres(config.Config.Database)
	if err != nil {
		log.Fatalf("error while loading postgreSQL: %s:", err)
	}

	database.Cache, err = database.SetupRedis(config.Config.Database)

	if err != nil {
		log.Fatalf("error while loading redis: %s:", err)
	}

	defer f.Close()

	log.SetOutput(f)

	// setting up web server middlewares
	middlewareManager := negroni.New(
		negroni.HandlerFunc(auth.IsAuthorized),
		negroni.HandlerFunc(auth.AccessPermission),
	)
	middlewareManager.Use(negroni.NewRecovery())
	middlewareManager.UseHandler(service.NewRouter())

	log.Println("Starting HTTP listener...")
	err = http.ListenAndServe(config.Config.ListenURL, middlewareManager)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Stop running application: %s", err)
}
