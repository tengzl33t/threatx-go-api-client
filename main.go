package main

import "C"
import (
	"log"
	"threatx-go-api-client/pkg"
	"time"
)

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}

func main() {
	defer timeTrack(time.Now())
	pkg.RunClient(
		"sites",
		"prod",
		"",
		nil,
		[]map[string]interface{}{{"command": "list", "customer_name": "soctest"}},
	)

}
