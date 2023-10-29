package main

import (
	"check-domain-api/internal/controllers"
	"check-domain-api/internal/service"
	"check-domain-api/internal/storage"
	"check-domain-api/pkg"
	"log"
	"net/http"
	"sync"
)

func main() {
	db, err := pkg.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	storage := storage.NewDomain(db)
	services := service.NewDomain(storage)
	handlers := controllers.NewHandler(services)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		handlers.Worker()
		wg.Done()
	}()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handlers.InitRouter(),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}
