package transport

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TelemetryDatum struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type TelemetryBatch struct {
	NodeID        string           `json:"node_id"`
	Data          []TelemetryDatum `json:"data"`
	TimestampUnix int64            `json:"timestamp_unix"`
}

type JSONSender struct {
	endpointURL string
	client      *http.Client
}

func NewJSONSender(endpointURL string) *JSONSender {
	return &JSONSender{
		endpointURL: endpointURL,
		client:      &http.Client{Timeout: 5 * time.Second},
	}
}

func (s *JSONSender) SendTelemetry(batch *TelemetryBatch) error {
	payload, err := json.Marshal(batch)
	if err != nil {
		return err
	}

	resp, err := s.client.Post(s.endpointURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		log.Printf("Failed to send telemetry: status %d\n", resp.StatusCode)
	}

	return nil
}