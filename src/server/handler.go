package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pasca-l/global-ip-retriever/network"
)

func handleIps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ips, err := network.RetrieveIps()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprintf("failed to retrieve IPs: %s", err),
		))
	}

	res, err := json.Marshal(ips)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprintf("failed to marshal data into json: %s", err),
		))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
