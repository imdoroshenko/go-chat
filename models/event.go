package models

import "time"

type Event struct {
	Client *Client `json:"-"`
	Channel string `json:"channel"`
	Type string `json:"type"`
	Value string `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

