package nats

import (
	models "L0/internal"
	"L0/internal/middleware"
	"errors"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"log"
)

type Nats struct {
	Conn    stan.Conn
	service models.UseCase
}

func NewBroker(useCase models.UseCase) (models.Nats, error) {
	nc, err := nats.Connect("http://nats:4222")
	if err != nil {

		return nil, errors.New("cannot connect nats")
	}
	defer nc.Close()

	sc, err := stan.Connect("test-cluster", "client",
		stan.NatsURL("http://nats:4222"))
	if err != nil {
		return nil, err
	}

	return &Nats{
		Conn:    sc,
		service: useCase,
	}, nil
}

func (n *Nats) Subscribe(handle func(msg *stan.Msg)) error {
	_, err := n.Conn.Subscribe("updates", handle, stan.DurableName("alex"))
	if err != nil {
		_ = n.Conn.Close()
		return err
	}
	return nil
}

func (n *Nats) Handler(msg *stan.Msg) {
	id, err := middleware.CheckModel(msg.Data)
	if err != nil {
		log.Println(err)
		return
	}

	err = n.service.Save(id, msg.Data)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Broker got new model")
}
