package envs

import (
	"log"
	"os"
	"strconv"
)

const (
	WebTrackerAppEnvKey                  = "WEB_TRACKER_APP_ENV"
	WebTrackerServerAddressKey           = "WEB_TRACKER_SERVER_ADDRESS"
	WebTrackerServerPortKey              = "WEB_TRACKER_SERVER_PORT"
	WebTrackerServerReadTimeoutKey       = "WEB_TRACKER_SERVER_READ_TIMEOUT"
	WebTrackerServerWriteTimeoutKey      = "WEB_TRACKER_SERVER_WRITE_TIMEOUT"
	WebTrackerServerIdleTimeoutKey       = "WEB_TRACKER_SERVER_IDLE_TIMEOUT"
	WebTrackerBloomExpectedElementsKey   = "WEB_TRACKER_BLOOM_EXPECTED_ELEMENTS"
	WebTrackerBloomFalsePositiveRateKey  = "WEB_TRACKER_BLOOM_FALSE_POSITIVE_RATE"
)

type Env struct {
	AppEnv                  string
	ServerAddress           string
	ServerPort              int
	ServerReadTimeout       int
	ServerWriteTimeout      int
	ServerIdleTimeout       int
	BloomExpectedElements   int
	BloomFalsePositiveRate  float64
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
		AppEnv:                  appEnv,
		ServerAddress:           serverAddress,
		ServerPort:              port,
		ServerReadTimeout:       getEnvAsInt(WebTrackerServerReadTimeoutKey, 30),
		ServerWriteTimeout:      getEnvAsInt(WebTrackerServerWriteTimeoutKey, 30),
		ServerIdleTimeout:       getEnvAsInt(WebTrackerServerIdleTimeoutKey, 30),
		BloomExpectedElements:   getEnvAsInt(WebTrackerBloomExpectedElementsKey, 1000),
		BloomFalsePositiveRate:  getEnvAsFloat64(WebTrackerBloomFalsePositiveRateKey, 0.01),
	}
}

func getEnvAsInt(key string, defaultVal int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
		log.Printf("[ENV] Invalid value for %s: %s, using default: %d", key, value, defaultVal)
	}
	return defaultVal
}

func getEnvAsFloat64(key string, defaultVal float64) float64 {
	if value := os.Getenv(key); value != "" {
		if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
			return floatVal
		}
		log.Printf("[ENV] Invalid value for %s: %s, using default: %f", key, value, defaultVal)
	}
	return defaultVal
}
