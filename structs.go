package main

import "encoding/json"

type qq struct {
	Name   string `json:"name"`
	Potato int32  `json:"potato"`
}

type qqMessage struct {
	Message qq `json:"message"`
}

type baseMessage struct {
	MessageType int32           `json:"messageType"`
	Message     json.RawMessage `json:"message"`
}
