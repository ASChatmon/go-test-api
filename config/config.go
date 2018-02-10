package config

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"go-test-api/storage"
	"go-test-api/types"
)

// Config is a holder for the environmental configuration
// used to run an instance of the tsr-workflow-service.  It is accessed through
// a context.Context, which is provided to all queue workers and HTTP handlers.
type Config struct {
	Log         *types.Logger
	DatabaseUrl string
	Connection  *storage.DatabaseContext
}

// New reads the environment variables provided either
// by the OS or etcd and loads them into a new Config instance.
func New() Config {
	// Loggging
	logger := logrus.New()
	logger.Formatter = new(logrus.JSONFormatter)
	logger.Level = logrus.InfoLevel
	entry := logger.WithFields(logrus.Fields{"transaction_id": "N/A"})
	log := &types.Logger{Log: entry}

	// Database - Normally these values go in an env file and are saved in the config!!!!
	databaseHost := "127.0.0.1"
	databasePort := "3306"
	databaseName := "cpu_metrics"
	databaseUser := "root"
	databasePassword := ""
	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", databaseUser, databasePassword, databaseHost, databasePort, databaseName)

	return Config{
		Log:         log,
		DatabaseUrl: databaseUrl}

}
