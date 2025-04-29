package main

import "C"
import "threatx-go-api-client/pkg"

func main() {
	pkg.RunClient(
		"sites",
		"prod",
		"",
		nil,
		[]map[string]interface{}{{"command": "list", "customer_name": "soctest"}},
	)

}
