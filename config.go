package main

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type configType struct {
	tcpBind   string
	dirs      []string
	shortPoll time.Duration
	longPoll  time.Duration
}

var (
	config configType
)

func configFromEnv() configType {

	nodirs := []string{}

	config.tcpBind = getEnvString("HDDPROXY_TCP_BIND", "127.0.0.1:8080")

	sep := getEnvString("HDDPROXY_DIRS_SEP", ":")

	config.dirs = getEnvStrings("HDDPROXY_DIRS", nodirs, sep)

	sp := getEnvInt("HDDPROXY_SHORTPOLL", 1)
	lp := getEnvInt("HDDPROXY_LONGPOLL", 5)

	config.shortPoll = time.Duration(sp) * time.Second
	config.longPoll = time.Duration(lp) * time.Second

	return config
}

func getEnvString(varname, def string) string {

	if v := os.Getenv(varname); v != "" {
		return v
	}
	return def
}

func getEnvInt(varname string, def int) int {

	if v := os.Getenv(varname); v != "" {

		i, err := strconv.Atoi(v)
		if err != nil {
			return def
		}

		return i
	}
	return def
}

func getEnvStrings(varname string, def []string, sep string) []string {

	if v := os.Getenv(varname); v != "" {

		ar := strings.Split(v, sep)

		return ar
	}
	return def
}
