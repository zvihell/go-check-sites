package models

import "time"

type Domain struct {
	ID          int
	Domain      string
	Latency     time.Duration
	Available   bool
	Last_update time.Time
}
