package main

import (
	"contacts-list/internal/pkg/connector"
	"contacts-list/internal/pkg/mongo"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"contacts-list/internal/app/contact/create"
	"contacts-list/internal/app/contact/remove"
	"contacts-list/internal/app/contact/update"
	contactslist "contacts-list/internal/app/contacts-list"
)

func main() {
	config, err := parse()
	if err != nil {
		panic(err)
	}

	mongoConn := mongo.New("users contacts", config.Mongo)

	err = connector.Connect(config.Connector, mongoConn)
	if err != nil {
		panic(err)
	}

	createRequest := create.Setup(mongoConn.Collection("users_contacts"), config.Mongo.RequestsTimeout)
	updateRequest := update.Setup(mongoConn.Collection("users_contacts"), config.Mongo.RequestsTimeout)
	removeRequest := remove.Setup(mongoConn.Collection("users_contacts"), config.Mongo.RequestsTimeout)
	contactsListRequest := contactslist.Setup(mongoConn.Collection("users_contacts"), config.Mongo.RequestsTimeout)

	router := mux.NewRouter()

	router.Use(accessControlMiddleware)

	router.Handle("/contacts", contactsListRequest).Methods(http.MethodPost, http.MethodOptions)
	router.Handle("/contact", createRequest).Methods(http.MethodPost, http.MethodOptions)
	router.Handle("/contact", updateRequest).Methods(http.MethodPatch, http.MethodOptions)
	router.Handle("/contact", removeRequest).Methods(http.MethodDelete, http.MethodOptions)

	log.Println("starting http server")

	srv := &http.Server{
		Handler:      router,
		Addr:         config.HTTP.URI,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
