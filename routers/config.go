package routers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"test/image-service/utils"
)

func RunServer(){
	r := mux.NewRouter()

	baseURL := r.PathPrefix("/api/v1").Subrouter()
	baseURL.HandleFunc("/upload", utils.UploadFile).Methods("POST")
	images := http.StripPrefix("/images/", http.FileServer(http.Dir("./uploads/")))
	r.PathPrefix("/images/").Handler(images)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":10000", handler))
}
