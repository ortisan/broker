package main

import (
	"ortisan-broker/go-user-service/infrastructure/di"
)

func main() {
	r, err := di.ConfigRouters()
	if err != nil {
		panic(err)
	}
	r.Run(":8081")
}
