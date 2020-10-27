package main

import (
	"contacts-list/internal/pkg/connector"
	"contacts-list/internal/pkg/mongo"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	router.Handle("/{user}/contacts", contactsListRequest).Methods(http.MethodGet, http.MethodOptions)
	router.Handle("/{user}/contact", createRequest).Methods(http.MethodPost, http.MethodOptions)
	router.Handle("/{user}/contact/{id}", updateRequest).Methods(http.MethodPatch, http.MethodOptions)
	router.Handle("/{user}/contact/{id}", removeRequest).Methods(http.MethodDelete, http.MethodOptions)

	srv := &http.Server{
		Handler:      router,
		Addr:         config.HTTP.URI,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("starting http server")
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGABRT)

	<-stop
	log.Println("stopping server")

	err = srv.Shutdown(context.Background())
	if err != nil {
		log.Println(err)
	}

	log.Println("server stopped")
}

func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
