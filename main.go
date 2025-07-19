package main

import (
	"kumparan-tech-test/api"
	"kumparan-tech-test/config"
	"sync"
)

func main() {
	//Init Config
	config.SetConfig()

	var wg sync.WaitGroup
	wg.Add(1)

	//Start HTTP / REST Server
	go api.StartHttpServer()

	wg.Wait()
}
