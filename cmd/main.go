package main

import "avito-tech/internal/app"

func main() {
	if err := app.AppRun(); err != nil {
		panic(err)
	}
}
