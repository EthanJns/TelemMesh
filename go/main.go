package main

import (
	"fmt"
	"log"
	"telem_mesh/collector"
	"telem_mesh/transport"
	"time"
)

func main() {
	fmt.Println("Test")
	TELEM_ENDPOINT := "http://localhost:4000/api/telemetry"
	sender := transport.NewJSONSender(TELEM_ENDPOINT)

	for {
		cpu_met, err := collector.CollectCPUUsage()
		if err != nil {
			log.Printf("Error collecting CPU metricvs: %v", err)
			continue
		}

		batch := transport.TelemetryBatch {
			NodeID: "node-1",
			Data: []transport.TelemetryDatum{},
			TimestampUnix: time.Now().Unix(),
		}

		for _, metric := range cpu_met {
			batch.Data = append(batch.Data, transport.TelemetryDatum{
				Name: metric.Name,
				Value: metric.Value,
				Unit: metric.Unit,
			})
		}

		err = sender.SendTelemetry(&batch)

		if err != nil {
			log.Printf("Error sending telem data :%v", err)
		}

		time.Sleep(10 * time.Second)
	}
}