package main

import (
	"github.com/EclipIIse/WB_L2/develop/11_HTTPServer/config"
	simpleapp "github.com/EclipIIse/WB_L2/develop/11_HTTPServer/simpleApp"
	"log"
)

func main() {
	config := config.NewDefConfig()

	s := simpleapp.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
