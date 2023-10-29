build:
	docker-compose build 
run:
	docker-compose up 

migrate:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up