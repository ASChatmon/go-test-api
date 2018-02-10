package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-test-api/types"
	//	"sort"
	//	"strconv"
	//	"strings"
	//	"time"
)

type DatabaseContext struct {
	Connection *sql.DB
}

func NewConnection(url string) (*DatabaseContext, error) {
	context := &DatabaseContext{}

	db, err := sql.Open("mysql", url)
	if err != nil {
		return context, err
	}

	err = db.Ping()
	if err != nil {
		return context, err
	}

	context.Connection = db

	return context, nil
}

/*
Get Method for all memory data
@Return memory list
	error
*/
func (c *DatabaseContext) getMemory(logger *types.Logger) ([]*types.MemoryData, error) {
	memoryData := []*types.MemoryData{}

	// memory is a reserved sql term. using to show workaround.
	statement := "Select total, available,used, percent_used, active, inactive, wired, buffers, timestamp from `memory`"

	rows, err := c.Connection.Query(statement)
	if err != nil {
		logger.LogError("getMemory", "query", err.Error())
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		m := types.MemoryData{}
		err := rows.Scan(&m.Total, &m.Available, &m.Used, &m.UsedPercent, &m.Active, &m.Inactive, &m.Wired, &m.Buffers, &m.Timestamp)

		if err != nil {
			logger.LogError("getMemory", "scan row", err.Error())
			return nil, err
		}
		memoryData = append(memoryData, &m)
	}

	return memoryData, rows.Err()
}

func (c *DatabaseContext) GetMemoryByTimestamp(timestamp string, logger *types.Logger) (types.MemoryData, error) {
	m := types.MemoryData{}

	// memory is a reserved sql term. using to show workaround.
	statement := fmt.Sprintf(`SELECT total, available,used, percent_used, active, inactive, wired, buffers from memory
			WHERE timestamp = '%s'
			LIMIT 1`, timestamp)

	err := c.Connection.QueryRow(statement).Scan(&m.Total, &m.Available, &m.Used, &m.UsedPercent, &m.Active, &m.Inactive, &m.Wired, &m.Buffers)

	if err != nil {
		logger.LogError("getMemoryByTimestamp", "scan row", err.Error())
		return m, err
	}

	return m, nil
}

func (c *DatabaseContext) InsertMemory(mem types.MemoryData, time string, logger *types.Logger) error {
	_, err := c.Connection.Exec(`INSERT into memory
			(total, available, used, percent_used, active, inactive, wired, buffers, timestamp)
			VALUES
			(? , ?, ?, ?, ?, ?, ?, ?, ? ) `,
		mem.Total, mem.Available, mem.Used, mem.UsedPercent, mem.Active, mem.Inactive, mem.Wired, mem.Buffers, time)

	if err != nil {
		logger.LogError("insertMemory", "exec", err.Error())
		return err
	}

	return nil
}

func (c *DatabaseContext) getCPU(logger *types.Logger) ([]*types.CPUData, error) {
	cpuData := []*types.CPUData{}

	// cpu is a reserved sql term. using to show workaround.
	statement := "Select cpu, vender_id, family, model, stepping, physical_id, core_id, cores, model_name, mhz, cache_size, timestamp from `cpu`"

	rows, err := c.Connection.Query(statement)
	if err != nil {
		logger.LogError("getCPU", "query", err.Error())
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		cpu := types.CPUData{}
		err := rows.Scan(&cpu.CPU, &cpu.VendorID, &cpu.Family, &cpu.Model, &cpu.Stepping, &cpu.PhysicalID, &cpu.CoreID, &cpu.Cores, &cpu.ModelName, &cpu.Mhz, &cpu.CacheSize, &cpu.Timestamp)

		if err != nil {
			logger.LogError("getCPU", "scan row", err.Error())
			return nil, err
		}
		cpuData = append(cpuData, &cpu)
	}

	return cpuData, rows.Err()
}

func (c *DatabaseContext) GetCPUByTimestamp(timestamp string, logger *types.Logger) (types.CPUData, error) {
	cpu := types.CPUData{}

	// cpu is a reserved sql term. using to show workaround.
	statement := fmt.Sprintf(`SELECT cpu, vender_id, family, model, stepping, physical_id, core_id, cores, model_name, mhz, cache_size
			FROM cpu
			WHERE timestamp = '%s'
			LIMIT 1`, timestamp)

	err := c.Connection.QueryRow(statement).Scan(&cpu.CPU, &cpu.VendorID, &cpu.Family, &cpu.Model, &cpu.Stepping, &cpu.PhysicalID, &cpu.CoreID, &cpu.Cores, &cpu.ModelName, &cpu.Mhz, &cpu.CacheSize)

	if err != nil {
		logger.LogError("GetCPUByTimestamp", "scan row", err.Error())
		return cpu, err
	}

	return cpu, nil
}

func (c *DatabaseContext) InsertCPU(cpu types.CPUData, time string, logger *types.Logger) error {
	_, err := c.Connection.Exec(`INSERT INTO cpu
			(cpu, vender_id, family, model, stepping, physical_id, core_id, cores, model_name, mhz, cache_size, timestamp)
			VALUES
			( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ) `,
		cpu.CPU, cpu.VendorID, cpu.Family, cpu.Model, cpu.Stepping, cpu.PhysicalID, cpu.CoreID, cpu.Cores, cpu.ModelName, cpu.Mhz, cpu.CacheSize, time)

	if err != nil {
		logger.LogError("insertCPU", "exec", err.Error())
		return err
	}

	return nil
}

func (c *DatabaseContext) getDisk(logger *types.Logger) ([]*types.DiskData, error) {
	diskData := []*types.DiskData{}

	// go sql.DB handles reserved names very well.
	statement := "SELECT fstype, total, free, used, used_percent, inodes_total, inodes_used, inodes_free, inodes_used_percent, timestamp FROM disk"

	rows, err := c.Connection.Query(statement)
	if err != nil {
		logger.LogError("getDisk", "query", err.Error())
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		disk := types.DiskData{}
		err := rows.Scan(&disk.Fstype, &disk.Total, &disk.Free, &disk.Used, &disk.UsedPercent, &disk.InodesTotal, &disk.InodesUsed, &disk.InodesFree, &disk.InodesUsedPercent, &disk.Timestamp)

		if err != nil {
			logger.LogError("getDisk", "scan row", err.Error())
			return nil, err
		}
		diskData = append(diskData, &disk)
	}

	return diskData, rows.Err()
}

func (c *DatabaseContext) GetDiskByTimestamp(timestamp string, logger *types.Logger) (types.DiskData, error) {
	disk := types.DiskData{}

	// cpu is a reserved sql term. using to show workaround.
	statement := fmt.Sprintf(`SELECT fstype, total, free, used, used_percent, inodes_total, inodes_used, inodes_free, inodes_used_percent
			FROM disk
			WHERE timestamp = '%s'
			LIMIT 1`, timestamp)

	err := c.Connection.QueryRow(statement).Scan(&disk.Fstype, &disk.Total, &disk.Free, &disk.Used, &disk.UsedPercent, &disk.InodesTotal, &disk.InodesUsed, &disk.InodesFree, &disk.InodesUsedPercent)
	if err != nil {
		logger.LogError("GetDiskByTimestamp", "scan row", err.Error())
		return disk, err
	}

	return disk, nil
}

func (c *DatabaseContext) InsertDisk(disk types.DiskData, time string, logger *types.Logger) error {
	_, err := c.Connection.Exec(`INSERT INTO disk
			(fstype, total, free, used, used_percent, inodes_total, inodes_used, inodes_free, inodes_used_percent, timestamp)
			VALUES
			( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) `,
		disk.Fstype, disk.Total, disk.Free, disk.Used, disk.UsedPercent, disk.InodesTotal, disk.InodesUsed, disk.InodesFree, disk.InodesUsedPercent, time)

	if err != nil {
		logger.LogError("insertDisk", "exec", err.Error())
		return err
	}

	return nil
}

func (c *DatabaseContext) GetAggregatedData(disk types.DiskData, time string, logger *types.Logger) error {
	_, err := c.Connection.Exec(`INSERT INTO disk
			(fstype, total, free, used, used_percent, inodes_total, inodes_used, inodes_free, inodes_used_percent, timestamp)
			VALUES
			( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) `,
		disk.Fstype, disk.Total, disk.Free, disk.Used, disk.UsedPercent, disk.InodesTotal, disk.InodesUsed, disk.InodesFree, disk.InodesUsedPercent, time)

	if err != nil {
		logger.LogError("insertDisk", "exec", err.Error())
		return err
	}

	return nil
}
