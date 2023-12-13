package main

import (
	"goodbye/app"
	"os"
)

func main() {
	app.Run(os.Getenv("ENV"))
}
