start:
	go run main.go

swagger:
	docker compose build swagger-ui
	docker compose up swagger-ui -d