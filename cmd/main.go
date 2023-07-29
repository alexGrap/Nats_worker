package main

import (
	"L0/config"
	"L0/db/postgres"
	"L0/internal/delivery"
	"L0/internal/delivery/nats"
	"L0/internal/usecase"
	"log"
)

func main() {
	viperConf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	conf, err := config.ParseConfig(viperConf)
	if err != nil {
		log.Fatal(err)
	}
	useCase := usecase.InitUseCase(conf)
	newBroker, err := nats.NewBroker(useCase)
	if err != nil {
		log.Fatal(err)
	}
	err = newBroker.Subscribe(newBroker.Handler)
	if err != nil {
		log.Fatal(err)
	}
	postgres.InitPsqlDB(conf)
	server := delivery.Fabric(conf, useCase)
	err = server.Hearing()
	if err != nil {
		return
	}

}
