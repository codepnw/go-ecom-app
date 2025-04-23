run:
	go run main.go

docker-up:
	docker compose --env-file dev.env up