package main

import (
	"flag"
	"fmt"
	"net/http"
	"slo19/frontend-masters-project/internal/app"
	"slo19/frontend-masters-project/internal/routes"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "backend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	defer app.DB.Close()

	app.Logger.Println("We are running our app on port ", port)
	r := routes.SetupRoutes(app)
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
