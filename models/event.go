package models

import "time"

type Event struct {
	id string `json:"id"`
	Client *Client `json:"-"`
	Status int `json:"status"`
	Resource string `json:"resource"`
	Method string `json:"method"`
	Message *Message `json:"message"`
	Channel *Channel `json:"channel"`
	Timestamp time.Time `json:"timestamp"`
}
