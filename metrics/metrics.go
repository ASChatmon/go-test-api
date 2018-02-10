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
func getCurrMemory(connection *storage.DatabaseContext, timestamp string, logger *types.Logger) (*mem.VirtualMemoryStat, error) {
	m, err := mem.VirtualMemory()
	if err != nil {
		return m, err
	}
	// store memory
	err = connection.InsertMemory(*m, timestamp, logger)
	if err != nil {
		return m, err
	}

	return m, err
}

// get and save curr CPU stats
func getCurrCPU(connection *storage.DatabaseContext, timestamp string, logger *types.Logger) (cpu.InfoStat, error) {

	c, err := cpu.Info()
	if err != nil {
		return cpu.InfoStat{}, err
	}

	// store cpu
	err = connection.InsertCPU(c[0], timestamp, logger)
	if err != nil {
		return cpu.InfoStat{}, err
	}

	return c[0], err
}

// get and save curr disk stats
func getCurrDisk(connection *storage.DatabaseContext, timestamp string, logger *types.Logger) (disk.UsageStat, error) {

	d, err := disk.Usage("/")
	if err != nil {
		return disk.UsageStat{}, err
	}

	// store cpu
	err = connection.InsertDisk(*d, timestamp, logger)
	if err != nil {
		return disk.UsageStat{}, err
	}

	return *d, err
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
