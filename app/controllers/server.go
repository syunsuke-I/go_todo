package controllers

import (
	"lesson/config"
	"net/http"
)

func StartMainSever() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
