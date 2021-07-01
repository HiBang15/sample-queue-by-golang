package main

import "time"

type WorkRequest struct {
	Name  string        `json:"name"`
	Delay time.Duration `json:"delay"`
}
