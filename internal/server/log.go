package server

/*
	Basic Commit Log Package - Data Structure for keeping track of records
*/

import (
	"fmt"
	"sync"
)

type Record struct {
	// Expects log entries in JSON format
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

type Log struct {
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
	// Initializing Log with zero values
	return &Log{}
}

// Using pointer to log to update all instances of that log
// Method Header -> method receiver, params, return type(s)
func (log *Log) Append(record Record) (uint64, error) {
	log.mu.Lock()
	// Deferring unlock of mutex until record is appended
	defer log.mu.Unlock()
	record.Offset = uint64(len(log.records))
	log.records = append(log.records, record)
	return record.Offset, nil
}

func (log *Log) Read(offset uint64) (Record, error) {
	log.mu.Lock()
	// Deferring unlock of mutex until record is read
	defer log.mu.Unlock()
	if offset >= uint64(len(log.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return log.records[offset], nil
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")
