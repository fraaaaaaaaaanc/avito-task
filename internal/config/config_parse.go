package config

import (
	"flag"
	"os"
)

func ParseConfFlags() *Flags {
	objFlags := parseFlags()
	parseEnv(objFlags)
	return objFlags
}

func parseEnv(flags *Flags) {
	if host := os.Getenv("RUN_ADDRESS"); host != "" {
		flags.Set(host)
	}

	if secretKeyJwtToken := os.Getenv("SECRET_KEY_FOR_COOKIE_TOKEN"); secretKeyJwtToken != "" {
		flags.SecretKeyJWTToken = secretKeyJwtToken
	}

	if dataBaseURI := os.Getenv("DATABASE_URI"); dataBaseURI != "" {
		flags.DataBaseURI = dataBaseURI
	}

	if accrualSystemAddress := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); accrualSystemAddress != "" {
		flags.AccrualSystemAddress = accrualSystemAddress
	}

	if LogFilePath := os.Getenv("LOG_FILE"); LogFilePath != "" {
		flags.LogFilePath = LogFilePath
	}
}

func parseFlags() *Flags {
	objFlags := newFlags()

	flag.Var(&objFlags.HTTPServerHost, "a", "address and port to run server")
	flag.StringVar(&objFlags.DataBaseURI, "d", objFlags.DataBaseURI, "database connection address:host user "+
		"password dbname sslmode")
	flag.StringVar(&objFlags.AccrualSystemAddress, "r", objFlags.AccrualSystemAddress,
		"address of the accrual calculation system")
	flag.StringVar(&objFlags.LogFilePath, "lf", objFlags.LogFilePath, "the path for the file to which the "+
		"logs will be written")
	flag.StringVar(&objFlags.ProjLvl, "ll", objFlags.ProjLvl, "project development stage")

	return &objFlags
}
