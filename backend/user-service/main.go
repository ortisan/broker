package main

import (
	"monolith/infrastructure/di"
)

func main() {
	r, err := di.InitializeRouters()
	if err != nil {
		panic(err)
	}
	r.Run(":8081")
}
