package main

import (
	"encoding/json"
	"github.com/SashaCrofter/network"
	"strconv"
	"time"
)

type Snapshot struct {
	ID               string // ID of the relevant network
	Timestamp        int64  // Timestamp (UNIX seconds) of the snapshot
	*network.Network        // Network data
}

// Creates a Snapshot from data that can be found in an HTML form. The
// timestamp is assumed to be an integer representing the UNIX time in
// seconds at which the snapshot was taken. If the timestamp is not
// supplied, the current UNIX time will be used.
func BuildSnapshot(id, timestamp, data string) (s Snapshot, err error) {
	n := &network.Network{}
	err = json.Unmarshal([]byte(data), n)
	if err != nil {
		return
	}

	unixtime, err := strconv.ParseInt(timestamp, 0, 64)
	if err != nil {
		l.Println("Could not parse timestamp; using current time instead")
		unixtime = time.Now().Unix()
		err = nil
	}
	s = Snapshot{
		ID:        id,
		Timestamp: unixtime,
		Network:   n,
	}
	return
}

// Stores a Snapshot in the database.
func Store(s Snapshot) (err error) {
	l.Printf("Pretending to store snapshot of %s\n", s.ID)
	return nil
}
