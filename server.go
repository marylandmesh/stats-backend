package main

import (
	"net/http"
)

func serve(iface, port string) (err error) {
	l.Printf("Starting server at %q on %s:%s", *fRoot, iface, port)
	http.HandleFunc(*fRoot+"pickup", handlePickup)
	http.HandleFunc(*fRoot+"delivery", handleDelivery)
	return http.ListenAndServe(iface+":"+port, nil)
}

func handlePickup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pickup, gov'ner?"))
}

func handleDelivery(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delivery, gov'ner?"))
}
