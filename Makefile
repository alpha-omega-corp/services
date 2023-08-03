db_create:
	cd config/envs/ && docker-compose up -d

db_init:
	go run migrations/main.go db init

db_reset:
	go run migrations/main.go db reset
