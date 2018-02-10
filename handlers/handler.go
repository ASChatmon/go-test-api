package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/zenazn/goji/web"
	"go-test-api/config"
	metrics "go-test-api/metrics"
	"net/http"
	"time"
)

func WorkerMetrics(config *config.Config) {
	_, err := metrics.GetCurrData(config.Connection, config.Log)
	if err != nil {
		config.Log.LogError("WorkerMetrics", "Execution error", err.Error())
	}
}

func GetCurrentMetics(config *config.Config, c web.C, w http.ResponseWriter, r *http.Request) {
	startNanos := time.Now().UnixNano()
	w.Header().Set("Content-Type", "application/json")

	metrics, err := metrics.GetCurrData(config.Connection, config.Log)
	if err != nil {
		config.Log.LogError("GetCurrentMetrics", "Calling API", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(metrics)
	if err != nil {
		config.Log.LogError("GetCurrentMetrics", "Unmarshal Systems", fmt.Sprintf("Failed to marshal JSON response. [Error: %s]", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	finishMilis := (time.Now().UnixNano() - startNanos) / 1000000

	config.Log.LogInfo("GetCurrentMetics", "Process Time", fmt.Sprintf("%d ms", finishMilis))
	w.Write(json)
}

func GetMeticsByTimestamp(config *config.Config, c web.C, w http.ResponseWriter, r *http.Request) {
	startNanos := time.Now().UnixNano()
	w.Header().Set("Content-Type", "application/json")

	timestamp := c.URLParams["timestamp"]

	metrics, err := metrics.GetDataByTimestamp(timestamp, config.Connection, config.Log)
	if err != nil {
		config.Log.LogError("GetMeticsByTimestamp", "Calling API", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(metrics)
	if err != nil {
		config.Log.LogError("GetMeticsByTimestamp", "Unmarshal Systems", fmt.Sprintf("Failed to marshal JSON response. [Error: %s]", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	finishMilis := (time.Now().UnixNano() - startNanos) / 1000000

	config.Log.LogInfo("GetCurrentMetics", "Process Time", fmt.Sprintf("%d ms", finishMilis))
	w.Write(json)
}

func GetAggregatedMetrics(config *config.Config, c web.C, w http.ResponseWriter, r *http.Request) {
	startNanos := time.Now().UnixNano()
	w.Header().Set("Content-Type", "application/json")

	metrics, err := metrics.GetAggregatedData(config.Connection, config.Log)
	if err != nil {
		config.Log.LogError("GetAggregatedMetrics", "Calling API", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(metrics)
	if err != nil {
		config.Log.LogError("GetAggregatedMetrics", "Unmarshal Systems", fmt.Sprintf("Failed to marshal JSON response. [Error: %s]", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	finishMilis := (time.Now().UnixNano() - startNanos) / 1000000

	config.Log.LogInfo("GetCurrentMetics", "Process Time", fmt.Sprintf("%d ms", finishMilis))
	w.Write(json)
}

func GetAverageMetrics(config *config.Config, c web.C, w http.ResponseWriter, r *http.Request) {
	startNanos := time.Now().UnixNano()
	w.Header().Set("Content-Type", "application/json")

	metrics, err := metrics.GetAverageData(config.Connection, config.Log)
	if err != nil {
		config.Log.LogError("GetAverageMetrics", "Calling API", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(metrics)
	if err != nil {
		config.Log.LogError("GetAverageMetrics", "Unmarshal Systems", fmt.Sprintf("Failed to marshal JSON response. [Error: %s]", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	finishMilis := (time.Now().UnixNano() - startNanos) / 1000000

	config.Log.LogInfo("GetCurrentMetics", "Process Time", fmt.Sprintf("%d ms", finishMilis))
	w.Write(json)
}
