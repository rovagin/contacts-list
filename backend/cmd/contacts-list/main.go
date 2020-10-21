package main

import (
	contactslist "contacts-list/internal/app/contacts-list"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	contactsListRequest := contactslist.Setup()

	router := mux.NewRouter()

	router.Handle("/contacts", contactsListRequest)

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
