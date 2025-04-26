package main

import (
	"flag"
	"fmt"
	"net/http"
	"slo19/frontend-masters-project/internal/app"
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

	app.Logger.Println("We are running our app on port ", port)

	http.HandleFunc("/health", HealthCheck)
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}
