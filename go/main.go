package main

import (
	"fmt"
	"log"
	"telem_mesh/collector"
	"telem_mesh/transport"
	"time"
)

func main() {
	fmt.Println("Starting telemetry collection...")
	TELEM_ENDPOINT := "http://localhost:4000/api/telemetry"
	sender := transport.NewJSONSender(TELEM_ENDPOINT)

	// Collect static CPU info once at startup
	fmt.Println("Collecting static CPU info...")
	cpuInfo, err := collector.CollectCPUInfo()
	if err != nil {
		log.Printf("Error collecting CPU info: %v", err)
	} else {
		batch := transport.TelemetryBatch{
			NodeID:        "node-1",
			Data:          []transport.TelemetryDatum{},
			TimestampUnix: time.Now().Unix(),
		}

		for _, metric := range cpuInfo {
			batch.Data = append(batch.Data, transport.TelemetryDatum{
				Name:  metric.Name,
				Value: metric.Value,
				Unit:  metric.Unit,
			})
		}

		err = sender.SendTelemetry(&batch)
		if err != nil {
			log.Printf("Error sending CPU info: %v", err)
		} else {
			fmt.Printf("Sent %d CPU info metrics\n", len(cpuInfo))
		}
	}

	// Main collection loop
	fmt.Println("Starting main collection loop...")
	for {
		// Option 1: Use your existing lightweight collection
		// cpu_met, err := collector.CollectCPUUsage()
		
		// Option 2: Use the new lightweight collection (CPU usage + per-core)
		cpu_met, err := collector.CollectLightweightCPUMetrics()
		
		// Option 3: Use comprehensive collection (includes CPU times)
		// cpu_met, err := collector.CollectAllCPUMetrics()

		if err != nil {
			log.Printf("Error collecting CPU metrics: %v", err)
			continue
		}

		batch := transport.TelemetryBatch{
			NodeID:        "node-1",
			Data:          []transport.TelemetryDatum{},
			TimestampUnix: time.Now().Unix(),
		}

		for _, metric := range cpu_met {
			batch.Data = append(batch.Data, transport.TelemetryDatum{
				Name:  metric.Name,
				Value: metric.Value,
				Unit:  metric.Unit,
			})
		}

		err = sender.SendTelemetry(&batch)
		if err != nil {
			log.Printf("Error sending telemetry data: %v", err)
		} else {
			fmt.Printf("Sent %d metrics at %s\n", len(cpu_met), time.Now().Format("15:04:05"))
		}

		time.Sleep(10 * time.Second)
	}
}