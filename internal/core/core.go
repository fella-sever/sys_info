package core

import (
	"fmt"
	"log"
)

func Start() {
	config, err := NewConfig("../configs/config.yaml")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(config)
}
