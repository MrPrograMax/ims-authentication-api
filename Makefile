build:
	docker-compose build ims-authentication-api

run:
	docker-compose up ims-authentication-api

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up


