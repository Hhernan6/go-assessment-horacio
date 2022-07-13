package main

import (
	"fmt"
	"go-assessment/internal/app"
	"log"
)

func main() {
	fmt.Println("Hello fellow golanger! Good luck!")

	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	app.Start()
}
