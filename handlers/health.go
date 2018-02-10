package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/zenazn/goji/web"
	"go-test-api/config"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Version string `json:"version"`
	Status  string `json:"status"`
}

func GetHealthHandler(config *config.Config, c web.C, w http.ResponseWriter, r *http.Request) {
	resp := Response{Status: "I'm Alive!"}
	w.Header().Set("Content-Type", "application/json")

	// read file
	b, err := ioutil.ReadFile("version.txt")
	if err != nil {
		resp.Version = fmt.Sprintf("Could not read version.txt: The error is [%+v]", err)
	} else {
		resp.Version = string(b)
	}

	json.NewEncoder(w).Encode(resp)
}
