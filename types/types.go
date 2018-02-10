package types

import (
	"github.com/Sirupsen/logrus"
)

// Logger is a wrapper around logrus that provides convenience
// methods for structured logs.
type Logger struct {
	Log *logrus.Entry
}


type Metrics struct {
	CPU    CPUData    `json:"cpu"`
	Memory MemoryData `json:"memory"`
	Disk   DiskData   `json:"disk"`
}

type CPUData struct {
	CPU        int32    `json:"cpu,omitempty"`
	VendorID   string   `json:"vendor_id,omitempty"`
	Family     string   `json:"family,omitempty"`
	Model      string   `json:"model,omitempty"`
	Stepping   int32    `json:"stepping,omitempty"`
	PhysicalID string   `json:"physical_id,omitempty"`
	CoreID     string   `json:"core_id,omitempty"`
	Cores      int32    `json:"cores,omitempty"`
	ModelName  string   `json:"model_name,omitempty"`
	Mhz        float64  `json:"mhz,omitempty"`
	CacheSize  int32    `json:"cache_size,omitempty"`
	Timestamp  string   `json:"timestamp"`
}

type MemoryData struct {
	Total uint64 `json:"total,omitempty"`
	Available uint64 `json:"available,omitempty"`
	Used uint64 `json:"used,omitempty"`
	UsedPercent float64 `json:"percent_used,omitempty"`
	Free uint64 `json:"free,omitempty"`
	Active   uint64 `json:"active,omitempty"`
	Inactive uint64 `json:"inactive,omitempty"`
	Wired    uint64 `json:"wired,omitempty"`
	Buffers      uint64 `json:"buffers,omitempty"`
	Cached       uint64 `json:"cached,omitempty"`
	Timestamp  string   `json:"timestamp,omitempty"`
}

type DiskData struct {
	Fstype            string  `json:"fstype,omitempty"`
	Total             uint64  `json:"total,omitempty"`
	Free              uint64  `json:"free,omitempty"`
	Used              uint64  `json:"used,omitempty"`
	UsedPercent       float64 `json:"used_percent,omitempty"`
	InodesTotal       uint64  `json:"inodes_total,omitempty"`
	InodesUsed        uint64  `json:"inodes_used,omitempty"`
	InodesFree        uint64  `json:"inodes_free,omitempty"`
	InodesUsedPercent float64 `json:"inodes_used_percent,omitempty"`
	Timestamp  	  string   `json:"timestamp,omitempty"`
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
