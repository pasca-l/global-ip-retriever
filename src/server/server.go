package server

import (
	"net/http"
)

func Serve() error {
	http.HandleFunc("/", handleIps)

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
