package config

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
	"time"
)

const (
	envProd    = "prod"
	envStaging = "staging"
)

type Config struct {
	DBConfig *DBConfig
	Metrics  *Metrics
	Log      *Log
	Env      string
}

type Log struct {
	Level     slog.Level
	Dest      io.Writer
	AddSource bool
}

type Client struct {
	Address string
	Timeout time.Duration
}

type DBConfig struct {
	Conn string
}

type Metrics struct {
	Address   string
	Subsystem string
	Namespace string
}

func MustNew() Config {
	return Config{
		DBConfig: &DBConfig{
			Conn: mustGetEnv("PG_CONN_STRING"),
		},
		Metrics: &Metrics{
			Address:   mustGetEnv("METRICS_ADDRESS"),
			Namespace: mustGetEnv("METRICS_NAMESPACE"),
			Subsystem: mustGetEnv("METRICS_SUBSYSTEM"),
		},
		Log: &Log{
			Level:     mustGetLogLevelFromEnv("LOG_LEVEL"),
			Dest:      mustGetDestFromEnv("LOG_DEST"),
			AddSource: mustGetBoolFromEnv("LOG_ADD_SOURCE"),
		},
		Env: os.Getenv("APP_ENVIRONMENT"),
	}
}

func (c *Config) IsProd() bool {
	return c.Env == envProd
}

func (c *Config) IsStaging() bool {
	return c.Env == envStaging
}

func mustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("env variable %s must be set", key))
	}
	return v
}

func mustGetDurationFromEnv(key string) time.Duration {
	s := mustGetEnv(key)

	v, err := time.ParseDuration(s)
	if err != nil {
		panic(fmt.Sprintf("'%v' value is not a duration", key))
	}

	return v
}

func mustGetBoolFromEnv(key string) bool {
	s := mustGetEnv(key)

	v, err := strconv.ParseBool(s)
	if err != nil {
		panic(fmt.Sprintf("'%v' value is not a boolean", key))
	}

	return v
}

func mustGetLogLevelFromEnv(key string) slog.Level {
	s := mustGetEnv(key)
	var level slog.Level

	switch s {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		panic(fmt.Sprintf("cannot resolve %s into clog.Level", key))
	}

	return level
}

func mustGetDestFromEnv(key string) io.Writer {
	s := mustGetEnv(key)

	switch s {
	case "stdout":
		return os.Stdout
	case "stderr":
		return os.Stderr
	default:
		return os.Stdout
	}
}
