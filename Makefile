migrate-up:
	migrate -path external/db/migration -database "mysql://root:root@tcp(localhost:3306)/simple_product" -verbose up

migrate-down:
	migrate -path external/db/migration -database "mysql://root:root@tcp(localhost:3306)/simple_product" -verbose down

start:
	go run main/app.go

run-db:
	docker pull mysql/mysql-server:latest
	docker run --name simple-product-db -e MYSQL_USER=root -e MYSQL_DATABASE=simple_product -e MYSQL_ROOT_PASSWORD=root -e MYSQL_PASSWORD=root -p 3306:3306 -d mysql/mysql-server:latest

run-redis:
	docker run --name simple-product-redis -d -p 6379:6379 redis redis-server --requirepass "root"