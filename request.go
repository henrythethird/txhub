package main

import "encoding/json"

type TransactionRequest struct {
	Raw  json.RawMessage `json:"raw"`
	Meta json.RawMessage `json:"meta"`
}
