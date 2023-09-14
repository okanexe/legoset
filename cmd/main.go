package main

import (
	"database/sql"
	"legoapi/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"legoapi/internal/handler"
	"legoapi/internal/repository"

	_ "github.com/lib/pq"
)

func main() {
	// PostgreSQL veritabanına bağlanma
	db, err := sql.Open("postgres", "user=okantest password=123456 dbname=lego sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Uygulama servisi ve HTTP işleyiciyi oluşturma
	legoSetRepository := repository.NewLegoSetRepository(db)
	legoSetService := service.NewLegoSetService(legoSetRepository)
	legoSetHandler := handler.NewLegoSetHandler(legoSetService)

	// HTTP yönlendiriciyi oluşturma
	router := mux.NewRouter()
	legoSetHandler.RegisterRoutes(router)

	// HTTP sunucusunu başlatma
	log.Println("HTTP sunucusu :8080 portunda çalışıyor...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
