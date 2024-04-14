package main

import (
	"avito-tech/internal/app"
	"log"
)

func main() {
	if err := app.AppRun(); err != nil {
		log.Panic(err.Error())
		panic(err)
	}
}
