package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func CreateMockContext() *context.Context {

	// Mock logger
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	entry := logger.WithFields(logrus.Fields{"transaction_id": "N/A"})
	log := &types.Logger{Log: entry}

	// Mock Config
	conf := config.New()
	client := statsd.NewStatsdClient(conf.StatsDHost, conf.StatsDPrefix)
	client.CreateSocket()
	collector := statsd.NewStatsdBuffer(time.Second*15, client)

	// Create our context
	context := &context.Context{
		Config: &conf,
		DB:     nil,
		Log:    log,
		Stats:  collector,
	}

	return context
}
