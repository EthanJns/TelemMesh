package collector

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

type Metric struct {
	Name string
	Value float64
	Unit string
}

func CollectCPUUsage() ([]Metric, error) {
	percentages, err := cpu.Percent(time.Second, false)

	if err != nil {
		return nil, err
	}

	return []Metric {
		{Name: "cpu_usage",
		Value: percentages[0],
		Unit: "percent",},
	}, nil
}