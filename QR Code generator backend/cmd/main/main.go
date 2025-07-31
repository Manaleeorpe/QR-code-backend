package main

import (
    "net/http"
    "qr-code-generator/pkg/config"
    "qr-code-generator/pkg/models"
    "qr-code-generator/pkg/routes"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

var handler http.Handler

func init() {
    config.Connect()
    db := config.GetDB()
    models.SetDB(db)
    db.AutoMigrate(&models.QRCode{})

    router := mux.NewRouter()
    routes.RegisterQRCodeGeneratorstoreRoutes(router)

    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowCredentials: true,
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"*"},
    })
    handler = c.Handler(router)
}

func Handler(w http.ResponseWriter, r *http.Request) {
    handler.ServeHTTP(w, r)
}
