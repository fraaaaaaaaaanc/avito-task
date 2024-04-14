package config

import (
	"errors"
	"strconv"
	"strings"
)

const (
	hostAddr    = "localhost"
	hostPort    = 8080
	addrDB      = "host=localhost password=1234 dbname=gofermart user=postgres sslmode=disable"
	accrualSA   = "http://localhost:8080"
	logLvlLocal = "local"
	redisURL = "chaches:6379"
)

type Flags struct {
	SecretKeyJWTToken    string
	AccrualSystemAddress string
	LogFilePath          string
	ProjLvl              string
	DataBaseURI          string
	RedisURL string
	HTTPServerHost
}

type HTTPServerHost struct {
	hostAddress string
	hostPort    int
}

func newFlags() Flags {
	return Flags{
		HTTPServerHost: HTTPServerHost{
			hostAddress: hostAddr,
			hostPort:    hostPort,
		},
		AccrualSystemAddress: accrualSA,
		DataBaseURI:          addrDB,
		ProjLvl:              logLvlLocal,
	}
}

func (hs *HTTPServerHost) String() string {
	return hs.hostAddress + ":" + strconv.Itoa(hs.hostPort)
}

func (hs *HTTPServerHost) Set(address string) error {
	if address == "" {
		hs.hostAddress = hostAddr
		hs.hostPort = hostPort
		return nil
	}
	listAddress := strings.Split(address, ":")
	if len(listAddress) != 2 {
		return errors.New("need address in a form host:port")
	}
	port, err := strconv.Atoi(listAddress[1])
	if err != nil {
		return err
	}
	hs.hostAddress = listAddress[0]
	hs.hostPort = port
	return nil
}
