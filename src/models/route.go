package models

import (
	"time"
)

type Route struct {
	ID       string    `json:"id"`
	Origin   string    `json:"origin"`
	Destination string `json:"destination"`
	Distance  float64   `json:"distance"`
	Duration  time.Duration `json:"duration"`
}