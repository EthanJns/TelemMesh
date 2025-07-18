package collector

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

type CpuMetrics struct {
	TimeStamp time.Time
	TotalUsage float64
	PerCore []float64
	Times []cpu.TimesStat
	LogicalCores int
	PhysicalCores int
}

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

func CPUUsageByCore()([]Metric, error) {
	percentages, err := cpu.Percent(time.Second, true)
	if err != nil {
		return nil, err
	}

	var metrics []Metric
	for i, percent := range percentages {
		metrics = append(metrics, Metric{
			Name:  fmt.Sprintf("cpu_core_%d_usage", i+1),
			Value: percent,
			Unit:  "percent",
		})
	}

	return metrics, nil
}

func CollectCPUTimes() ([]Metric, error) {
	times, err := cpu.Times(false) // false = aggregate across all CPUs
	if err != nil {
		return nil, err
	}

	if len(times) == 0 {
		return nil, fmt.Errorf("no CPU time data available")
	}

	cpuTime := times[0] // Get the first (aggregate) entry
	return []Metric{
		{Name: "cpu_time_user", Value: cpuTime.User, Unit: "seconds"},
		{Name: "cpu_time_system", Value: cpuTime.System, Unit: "seconds"},
		{Name: "cpu_time_idle", Value: cpuTime.Idle, Unit: "seconds"},
		{Name: "cpu_time_nice", Value: cpuTime.Nice, Unit: "seconds"},
		{Name: "cpu_time_iowait", Value: cpuTime.Iowait, Unit: "seconds"},
		{Name: "cpu_time_irq", Value: cpuTime.Irq, Unit: "seconds"},
		{Name: "cpu_time_softirq", Value: cpuTime.Softirq, Unit: "seconds"},
		{Name: "cpu_time_steal", Value: cpuTime.Steal, Unit: "seconds"},
		{Name: "cpu_time_guest", Value: cpuTime.Guest, Unit: "seconds"},
		{Name: "cpu_time_guest_nice", Value: cpuTime.GuestNice, Unit: "seconds"},
	}, nil
}

// CollectCPUTimesPerCore - get CPU time breakdown per core
func CollectCPUTimesPerCore() ([]Metric, error) {
	times, err := cpu.Times(true) // true = per CPU core
	if err != nil {
		return nil, err
	}

	var metrics []Metric
	for i, cpuTime := range times {
		metrics = append(metrics, []Metric{
			{Name: fmt.Sprintf("cpu_core_%d_time_user", i+1), Value: cpuTime.User, Unit: "seconds"},
			{Name: fmt.Sprintf("cpu_core_%d_time_system", i+1), Value: cpuTime.System, Unit: "seconds"},
			{Name: fmt.Sprintf("cpu_core_%d_time_idle", i+1), Value: cpuTime.Idle, Unit: "seconds"},
			{Name: fmt.Sprintf("cpu_core_%d_time_iowait", i+1), Value: cpuTime.Iowait, Unit: "seconds"},
		}...)
	}

	return metrics, nil
}

// CollectCPUInfo - static CPU information (call once at startup)
func CollectCPUInfo() ([]Metric, error) {
	info, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	logicalCores, err := cpu.Counts(true)
	if err != nil {
		return nil, err
	}

	physicalCores, err := cpu.Counts(false)
	if err != nil {
		return nil, err
	}

	var metrics []Metric
	if len(info) > 0 {
		metrics = append(metrics, []Metric{
			{Name: "cpu_cores_logical", Value: float64(logicalCores), Unit: "count"},
			{Name: "cpu_cores_physical", Value: float64(physicalCores), Unit: "count"},
			{Name: "cpu_mhz", Value: info[0].Mhz, Unit: "mhz"},
			{Name: "cpu_cache_size", Value: float64(info[0].CacheSize), Unit: "kb"},
		}...)
	}

	return metrics, nil
}

// CollectAllCPUMetrics - convenience function to collect all CPU metrics
func CollectAllCPUMetrics() ([]Metric, error) {
	var allMetrics []Metric

	// Basic CPU usage
	cpuUsage, err := CollectCPUUsage()
	if err != nil {
		return nil, fmt.Errorf("failed to collect CPU usage: %v", err)
	}
	allMetrics = append(allMetrics, cpuUsage...)

	// Per-core usage
	perCoreUsage, err := CPUUsageByCore()
	if err != nil {
		return nil, fmt.Errorf("failed to collect per-core CPU usage: %v", err)
	}
	allMetrics = append(allMetrics, perCoreUsage...)

	// CPU times
	cpuTimes, err := CollectCPUTimes()
	if err != nil {
		return nil, fmt.Errorf("failed to collect CPU times: %v", err)
	}
	allMetrics = append(allMetrics, cpuTimes...)

	return allMetrics, nil
}


func CollectLightweightCPUMetrics() ([]Metric, error) {
	var allMetrics []Metric

	cpuUsage, err := CollectCPUUsage()
	if err != nil {
		return nil, fmt.Errorf("failed to collect CPU usage: %v", err)
	}
	allMetrics = append(allMetrics, cpuUsage...)

	perCoreUsage, err := CPUUsageByCore()
	if err != nil {
		return nil, fmt.Errorf("failed to collect per-core CPU usage: %v", err)
	}
	allMetrics = append(allMetrics, perCoreUsage...)

	return allMetrics, nil
}