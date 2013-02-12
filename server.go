package main

import (
	"errors"
	"fmt"
	"net/http"
)

// Server errors, both to be reported to the admin and to the
// connecting user.
var (
	ErrNoID   = errors.New("no id presented")
	ErrNoData = errors.New("no data field presented")

	ErrInvalidData = errors.New("data could not be parsed: %s")
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
		l.Printf("Delivery from %s aborted: %s\n", r.RemoteAddr, ErrNoID)
		http.Error(w, ErrNoID.Error(), http.StatusBadRequest)
		return
	}
	// Now retrieve the data, and abort if it's not given.
	data := r.FormValue("data")
	if len(data) == 0 {
		l.Printf("Delivery from %s aborted: %s\n", r.RemoteAddr, ErrNoData)
		http.Error(w, ErrNoData.Error(), http.StatusBadRequest)
		return
	}

	// BUG(DuoNoxSol): There is no authentication. This should be
	// added. The general idea is that the ID is looked up and matched
	// with a password. Then, the hash (probably SHA-256) of the data
	// and id are matched with one provided in the request. If they
	// don't match, abort.

	// Try to create a snapshot from the given information, and report
	// if it fails.
	s, err := BuildSnapshot(id, r.FormValue("timestamp"), data)
	if err != nil {
		report := fmt.Errorf(ErrInvalidData.Error(), err)
		l.Printf("Delivery from %s aborted: %s\n", r.RemoteAddr,
			report)
		http.Error(w, report.Error(), http.StatusBadRequest)
		return
	}

	// Finally, register it in the database. If this fails, report it
	// both in the log and to the web.
	err = Store(s)
	if err != nil {
		l.Printf("Attempt to store snapshot failed: %s\n", err)
		http.Error(w, "Database store failed.",
			http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Successful."))
}
