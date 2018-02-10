package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-test-api/types"
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

//get current memory stats
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

// get memory from db by timestamp
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

// get aggregated memory data
func (c *DatabaseContext) GetMemoryAggregates(logger *types.Logger) (types.MemoryData, error) {
	mems, err := c.getMemory(logger)
	if err != nil {
		return types.MemoryData{}, err
	}
	//this is faster in db. just showing some go syntax

	aggM := types.MemoryData{Total: 0,
		Free:        0,
		Available:   0,
		Used:        0,
		UsedPercent: 0,
		Active:      0,
		Inactive:    0,
		Wired:       0,
		Buffers:     0,
		Cached:      0}
	//this is faster in db. just showing some go syntax
	for _, value := range mems {
		aggM.Total += value.Total
		aggM.Free += value.Free
		aggM.Available += value.Available
		aggM.Used += value.Used
		aggM.UsedPercent += value.UsedPercent
		aggM.Active += value.Active
		aggM.Inactive += value.Inactive
		aggM.Wired += value.Wired
		aggM.Buffers += value.Buffers
		aggM.Cached += value.Cached
	}

	return aggM, nil
}

// get average memory data
func (c *DatabaseContext) GetMemoryAverages(logger *types.Logger) (types.MemoryData, error) {
	m := types.MemoryData{}

	// memory is a reserved sql term. using to show workaround.
	statement := "SELECT FLOOR(AVG(total)), FLOOR(AVG(available)), FLOOR(AVG(used)), AVG(percent_used), FLOOR(AVG(active)), FLOOR(AVG(inactive)), FLOOR(AVG(wired)), FLOOR(AVG(buffers)) from memory"

	err := c.Connection.QueryRow(statement).Scan(&m.Total, &m.Available, &m.Used, &m.UsedPercent, &m.Active, &m.Inactive, &m.Wired, &m.Buffers)

	if err != nil {
		logger.LogError("GetMemoryAverages", "scan row", err.Error())
		return m, err
	}

	return m, nil
}

// insert memory data into db
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

// get current cpu stats
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

// get cpu from db by timestamp
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

// get aggregated cpu data
func (c *DatabaseContext) GetCPUAggregates(logger *types.Logger) (types.CPUData, error) {
	cpu, err := c.getCPU(logger)
	if err != nil {
		return types.CPUData{}, err
	}

	//this is faster in db. just showing some go syntax
	aggCPU := types.CPUData{Stepping: 0,
		CacheSize: 0,
		Cores:     0}
	for _, value := range cpu {
		aggCPU.Stepping += value.Stepping
		aggCPU.CacheSize += value.CacheSize
		aggCPU.Cores += value.Cores
	}

	return aggCPU, nil
}

// get averaged cpu data
func (c *DatabaseContext) GetCPUAverages(logger *types.Logger) (types.CPUData, error) {
	cpu := types.CPUData{}

	// cpu is a reserved sql term. using to show workaround.
	statement := "SELECT FLOOR(AVG(stepping)), FLOOR(AVG(cores)), FLOOR(AVG(mhz)), FLOOR(AVG(cache_size)) FROM cpu"

	err := c.Connection.QueryRow(statement).Scan(&cpu.Stepping, &cpu.Cores, &cpu.Mhz, &cpu.CacheSize)

	if err != nil {
		logger.LogError("GetCPUAverages", "scan row", err.Error())
		return cpu, err
	}

	return cpu, nil
}

// insert cpu data into db
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

// get current disk data
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

// get disk data by timestamp
func (c *DatabaseContext) GetDiskByTimestamp(timestamp string, logger *types.Logger) (types.DiskData, error) {
	disk := types.DiskData{}

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

// insert disk data into db
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

// get aggregated disk data
func (c *DatabaseContext) GetDiskAggregates(logger *types.Logger) (types.DiskData, error) {
	disks, err := c.getDisk(logger)
	if err != nil {
		return types.DiskData{}, err
	}
	//this is faster in db. just showing some go syntax

	aggD := types.DiskData{Total: 0,
		Free:              0,
		Used:              0,
		UsedPercent:       0,
		InodesTotal:       0,
		InodesFree:        0,
		InodesUsed:        0,
		InodesUsedPercent: 0}
	//this is faster in db. just showing some go syntax
	for _, value := range disks {
		aggD.Total += value.Total
		aggD.Free += value.Free
		aggD.Used += value.Used
		aggD.UsedPercent += value.UsedPercent
		aggD.InodesTotal += value.InodesTotal
		aggD.InodesFree += value.InodesFree
		aggD.InodesUsed += value.InodesUsed
		aggD.InodesUsedPercent += value.InodesUsedPercent
	}

	return aggD, nil
}

// get averaged disk data
func (c *DatabaseContext) GetDiskAverages(logger *types.Logger) (types.DiskData, error) {
	disk := types.DiskData{}
	// db is fast. let's try that
	statement := `SELECT FLOOR(AVG(total)), FLOOR(AVG(free)), FLOOR(AVG(used)), AVG(used_percent), FLOOR(AVG(inodes_total)), FLOOR(AVG(inodes_used)), FLOOR(AVG(inodes_free)), AVG(inodes_used_percent) FROM disk`

	err := c.Connection.QueryRow(statement).Scan(&disk.Total, &disk.Free, &disk.Used, &disk.UsedPercent, &disk.InodesTotal, &disk.InodesUsed, &disk.InodesFree, &disk.InodesUsedPercent)
	if err != nil {
		logger.LogError("GetDiskAverages", "scan row", err.Error())
		return disk, err
	}

	return disk, nil

}
