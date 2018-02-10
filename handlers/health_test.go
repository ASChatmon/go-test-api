package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/cbdr/tsr-workflow-service/config"
	"github.com/cbdr/tsr-workflow-service/context"
	"github.com/cbdr/tsr-workflow-service/types"
	"github.com/quipo/statsd"
	"github.com/stretchr/testify/assert"
	"github.com/zenazn/goji/web"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func MockContext() *context.Context {

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

func Test_Health_VersionError(t *testing.T) {

	ctx := MockContext()

	request, _ := http.NewRequest("GET", "", nil)
	recorder := httptest.NewRecorder()

	c := web.C{}

	GetHealthHandler(ctx, c, recorder, request)

	var r response
	_ = json.Unmarshal(recorder.Body.Bytes(), &r)

	assert.Equal(t, 200, recorder.Code, fmt.Sprintf("Status - Expected: %d, Actual: %d", 200, recorder.Code))
	assert.Equal(t, "I'm Alive!", r.Status, fmt.Sprintf("Body - Expected: %s, Actual: %s", "I'm Alive!", r.Status))
	assert.Contains(t, r.Version, "Could not read version.txt: ", fmt.Sprintf("Version Contains - Expected: %s, Actual: %s", "Could not read version.txt: ", r.Version))
}
