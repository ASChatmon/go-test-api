package handlers

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"go-test-api/storage"
	"go-test-api/types"
	"time"
)

// get and save curr memory stats
func getCurrMemory(connection *storage.DatabaseContext, timestamp string, logger *types.Logger) (*types.MemoryData, error) {
	m, err := mem.VirtualMemory()
	if err != nil {
		return &types.MemoryData{}, err
	}

	memory := types.MemoryData{
		Total:       m.Total,
		Available:   m.Available,
		Used:        m.Used,
		UsedPercent: m.UsedPercent,
		Free:        m.Free,
		Active:      m.Active,
		Inactive:    m.Inactive,
		Wired:       m.Wired,
		Buffers:     m.Buffers,
		Cached:      m.Cached,
	}
	// store memory
	err = connection.InsertMemory(memory, timestamp, logger)
	if err != nil {
		return &types.MemoryData{}, err
	}

	return &memory, err
}

// get and save curr CPU stats
func getCurrCPU(connection *storage.DatabaseContext, timestamp string, logger *types.Logger) (types.CPUData, error) {

	c, err := cpu.Info()
	if err != nil || len(c) == 0 {
		return types.CPUData{}, err
	}

	cpu := types.CPUData{
		CPU:        c[0].CPU,
		VendorID:   c[0].VendorID,
		Family:     c[0].Family,
		Model:      c[0].Model,
		Stepping:   c[0].Stepping,
		PhysicalID: c[0].PhysicalID,
		CoreID:     c[0].CoreID,
		Cores:      c[0].Cores,
		ModelName:  c[0].ModelName,
		Mhz:        c[0].Mhz,
		CacheSize:  c[0].CacheSize,
	}

	// store cpu
	err = connection.InsertCPU(cpu, timestamp, logger)
	if err != nil {
		return types.CPUData{}, err
	}

	return cpu, err
}

// get and save curr disk stats
func getCurrDisk(connection *storage.DatabaseContext, timestamp string, logger *types.Logger) (types.DiskData, error) {

	d, err := disk.Usage("/")
	if err != nil {
		return types.DiskData{}, err
	}

	disk := types.DiskData{
		Fstype:            d.Fstype,
		Total:             d.Total,
		Free:              d.Free,
		Used:              d.Used,
		UsedPercent:       d.UsedPercent,
		InodesTotal:       d.InodesTotal,
		InodesUsed:        d.InodesUsed,
		InodesFree:        d.InodesFree,
		InodesUsedPercent: d.InodesUsedPercent,
	}
	// store cpu
	err = connection.InsertDisk(disk, timestamp, logger)
	if err != nil {
		return types.DiskData{}, err
	}

	return disk, err
}

func GetCurrData(connection *storage.DatabaseContext, logger *types.Logger) (types.Metrics, error) {
	metrics := types.Metrics{}

	// be sure all metrics have same time for easy grabbing later
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	mem, err := getCurrMemory(connection, timestamp, logger)
	if err != nil {
		return metrics, err
	}

	metrics.Memory = *mem

	cpu, err := getCurrCPU(connection, timestamp, logger)
	if err != nil {
		return metrics, err
	}

	metrics.CPU = cpu

	disk, err := getCurrDisk(connection, timestamp, logger)
	if err != nil {
		return metrics, err
	}

	metrics.Disk = disk

	return metrics, nil
}

func GetDataByTimestamp(timestamp string, connection *storage.DatabaseContext, logger *types.Logger) (types.Metrics, error) {
	metrics := types.Metrics{}

	mem, err := connection.GetMemoryByTimestamp(timestamp, logger)
	if err != nil {
		return metrics, err
	}
	metrics.Memory = mem

	cpu, err := connection.GetCPUByTimestamp(timestamp, logger)
	if err != nil {
		return metrics, err
	}
	metrics.CPU = cpu

	disk, err := connection.GetDiskByTimestamp(timestamp, logger)
	if err != nil {
		return metrics, err
	}
	metrics.Disk = disk

	return metrics, nil
}

func GetAggregatedData(connection *storage.DatabaseContext, logger *types.Logger) (types.Metrics, error) {
	metrics := types.Metrics{}

	mem, err := connection.GetMemoryAggregates(logger)
	if err != nil {
		return metrics, err
	}
	metrics.Memory = mem

	cpu, err := connection.GetCPUAggregates(logger)
	if err != nil {
		return metrics, err
	}
	metrics.CPU = cpu

	disk, err := connection.GetDiskAggregates(logger)
	if err != nil {
		return metrics, err
	}
	metrics.Disk = disk

	return metrics, nil
}

func GetAverageData(connection *storage.DatabaseContext, logger *types.Logger) (types.Metrics, error) {
	metrics := types.Metrics{}

	mem, err := connection.GetMemoryAverages(logger)
	if err != nil {
		return metrics, err
	}
	metrics.Memory = mem

	cpu, err := connection.GetCPUAverages(logger)
	if err != nil {
		return metrics, err
	}
	metrics.CPU = cpu

	disk, err := connection.GetDiskAverages(logger)
	if err != nil {
		return metrics, err
	}
	metrics.Disk = disk

	return metrics, nil
}