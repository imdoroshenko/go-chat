package main

import "time"

type message struct {
	channel *appChannel
	client *client
	Message string `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
