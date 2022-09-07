package platform

import (
	"encoding/json"
	"time"
)

const (
	success = "0"
)

type Reply struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type JobStatusInfo struct {
	JobId  string    `json:"jobId"`
	Status string    `json:"status"`
	Time   time.Time `json:"time"`
}