build:
	docker compose build

run:
	docker compose up

make-migration:
	docker compose run --rm migrate create -ext sql -dir ./migrations -seq ${migration_name}
mm: make-migration

migrate:
	docker compose run --rm migrate -path=/migrations/ -database postgres://postgres:password@postgres:5432/go_test?sslmode=disable up
m: migrate

count ?= 1 # use -all to go all the way down
migrate-down:
	docker compose run --rm migrate -path=/migrations/ -database postgres://postgres:password@postgres:5432/go_test?sslmode=disable down ${count}
md: migrate-down

destroy:
	docker compose stop
	docker compose down -v --rmi local

# Examples of running kafka binaries that exist inside the container
topic ?= 
kafka-describe-topic:
	docker compose exec kafka ./opt/kafka/bin/kafka-topics.sh --describe --topic ${topic} --bootstrap-server kafka:9092

kafka-console-consumer:
	docker compose exec kafka ./opt/kafka/bin/kafka-console-consumer.sh --topic ${topic} --from-beginning --bootstrap-server kafka:9092
