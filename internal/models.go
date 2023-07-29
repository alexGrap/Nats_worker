package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/stan.go"
	"time"
)

type Repository interface {
	CheckIfExist() (map[string][]byte, error)
	Create(id string, jmodel []byte) error
	InitTable() error
}

type UseCase interface {
	Save(id string, jmodel []byte) error
	Get(id string) ([]byte, error)
	GetAll() ([]StoreStruct, error)
	Restore() error
}

type Rest interface {
	Hearing() error
	GetOne(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
}

type Nats interface {
	Subscribe(magicFunc func(msg *stan.Msg)) error
	Handler(msg *stan.Msg)
}

type Model struct {
	OrderUid    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
	Entry       string `json:"entry"`
	Delivery    struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction"`
		RequestId    string `json:"request_id"`
		Currency     string `json:"currency"`
		Provider     string `json:"provider"`
		Amount       int    `json:"amount"`
		PaymentDt    int    `json:"payment_dt"`
		Bank         string `json:"bank"`
		DeliveryCost int    `json:"delivery_cost"`
		GoodsTotal   int    `json:"goods_total"`
		CustomFee    int    `json:"custom_fee"`
	} `json:"payment"`
	Items []struct {
		ChrtId      int    `json:"chrt_id"`
		TrackNumber string `json:"track_number"`
		Price       int    `json:"price"`
		Rid         string `json:"rid"`
		Name        string `json:"name"`
		Sale        int    `json:"sale"`
		Size        string `json:"size"`
		TotalPrice  int    `json:"total_price"`
		NmId        int    `json:"nm_id"`
		Brand       string `json:"brand"`
		Status      int    `json:"status"`
	} `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

type StoreStruct struct {
	Id   string
	Body []byte
}
