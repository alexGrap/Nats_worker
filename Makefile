all:
	docker-compose up -d --build

nats:
	docker-compose up -d --build nats
server:
	docker-compose up -d --build wb_l0

publisher:
	go run script/main.go