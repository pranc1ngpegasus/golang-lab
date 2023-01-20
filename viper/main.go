package main

import (
	"log"
)

func main() {
	config, err := NewConfig()
	if err != nil {
		panic(err)
	}

	log.Default().Printf("%+v", config.Debug())
	log.Default().Printf("%+v", config.Server())
}
