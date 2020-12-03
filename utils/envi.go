package utils

import "github.com/h4lim/go-sdk/config"

func GetRunMode() string {
	serverMode := config.MustGetString("server.mode")
	return serverMode
}
