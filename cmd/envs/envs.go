package envs

import (
	"log"
	"os"
	"strconv"
)

const (
	WebTrackerAppEnvKey         = "WEB_TRACKER_APP_ENV"
	WebTrackerServerAddressKey  = "WEB_TRACKER_SERVER_ADDRESS"
	WebTrackerServerPortKey     = "WEB_TRACKER_SERVER_PORT"
)

type Env struct {
	AppEnv        string
	ServerAddress string
	ServerPort    int
}

func NewEnv() *Env {
	appEnv := os.Getenv(WebTrackerAppEnvKey)
	if appEnv == "" {
		log.Fatalf("[ENV] App environment not set")
	}

	serverAddress := os.Getenv(WebTrackerServerAddressKey)
	if serverAddress == "" {
		log.Fatalf("[ENV] Server address not set")
	}

	port, err := strconv.Atoi(os.Getenv(WebTrackerServerPortKey))
	if err != nil {
		log.Fatalf("[ENV] Invalid server port: %v", err)
	}

	return &Env{
		AppEnv:        os.Getenv(WebTrackerAppEnvKey),
		ServerAddress: os.Getenv(WebTrackerServerAddressKey),
		ServerPort:    port,
	}
}
