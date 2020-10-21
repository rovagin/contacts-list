package main

import (
	"contacts-list/internal/app/contact/create"
	getbyname "contacts-list/internal/app/contact/get-by-name"
	getbyphone "contacts-list/internal/app/contact/get-by-phone"
	"contacts-list/internal/app/contact/remove"
	"contacts-list/internal/app/contact/update"
	contactslist "contacts-list/internal/app/contacts-list"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	config, err := parse()
	if err != nil {
		panic(err)
	}

	createRequest := create.Setup()
	updateRequest := update.Setup()
	removeRequest := remove.Setup()
	getByNameRequest := getbyname.Setup()
	getByPhoneRequest := getbyphone.Setup()
	contactsListRequest := contactslist.Setup()

	router := mux.NewRouter()

	router.Handle("/contacts", contactsListRequest).Methods(http.MethodPost)
	router.Handle("/contact", createRequest).Methods(http.MethodPost)
	router.Handle("/contact/{id}", updateRequest).Methods(http.MethodPatch)
	router.Handle("/contact/{id}", removeRequest).Methods(http.MethodDelete)
	router.Handle("/contact/by-name", getByNameRequest).Methods(http.MethodPost)
	router.Handle("/contact/by-phone", getByPhoneRequest).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      router,
		Addr:         config.HTTP.URI,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
