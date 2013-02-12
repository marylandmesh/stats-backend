package main

import (
	"errors"
	"fmt"
	"net/http"
)

// Server errors, both to be reported to the admin and to the
// connecting user.
var (
	ErrNoID = errors.New("no id presented")
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
	// Retrieve the ID from the form, and abort if it's not given.
	id := r.FormValue("id")
	if len(id) == 0 {
		l.Printf("Delivery from %s aborted: %s", r.RemoteAddr, ErrNoID)
		http.Error(w, ErrNoID.Error(), http.StatusBadRequest)
		return
	}

	l.Printf("Getting delivery from %s, %q",
		r.RemoteAddr, r.FormValue("id"))
	fmt.Fprintf(w, "Got %s\n", r.FormValue("body"))
}
