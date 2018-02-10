package types

import (
	"github.com/Sirupsen/logrus"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// Logger is a wrapper around logrus that provides convenience
// methods for structured logs.
type Logger struct {
	Log *logrus.Entry
}

type Metrics struct {
	CPU    cpu.InfoStat          `json:"cpu"`
	Memory mem.VirtualMemoryStat `json:"memory"`
	Disk   disk.UsageStat        `json:"disk"`
}

/* this may seem odd but the merging of timestamp is less convoluted and clunky
than recreating the individual structs just to add time.
*/
type MetricsData struct {
	CPU    CPUData    `json:"cpu"`
	Memory MemoryData `json:"memory"`
	Disk   DiskData   `json:"disk"`
}

type Timestamp struct {
	Timestamp string `json:"timestamp"`
}

type CPUData struct {
	Timestamp
	cpu.InfoStat
}
type MemoryData struct {
	Timestamp
	mem.VirtualMemoryStat
}
type DiskData struct {
	Timestamp
	disk.UsageStat
}

// LogInfo logs a message at INFO level via logrus to STDOUT.
func (l *Logger) LogInfo(function string, action string, message string) {
	l.Log.WithFields(logrus.Fields{
		"function": function,
		"action":   action,
		"message":  message,
	}).Info("go-test-api")
}

// LogInfo logs a message at ERROR level via logrus to STDOUT.
func (l *Logger) LogError(function string, action string, err string) {
	l.Log.WithFields(logrus.Fields{
		"function": function,
		"action":   action,
		"error":    err,
	}).Error("go-test-api")
}
