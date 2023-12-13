include .env

create_container:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:12-alpine

stop_containers:
	@echo "Stopping other docker containers"
	@if [ $$(docker ps -q) ]; then \
		docker stop $$(docker ps -q); \
		echo "Docker containers stopped."; \
	else \
		echo "No active containers found."; \
	fi

create_db:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${USER} --owner=${USER} ${DB_NAME}

start_container:
	docker start ${DB_DOCKER_CONTAINER}

stop_container:
	docker stop ${DB_DOCKER_CONTAINER}

create_migrations:
	sqlx migrate add -r init

migrate_up:
	sqlx migrate run --database-url postgres://Wexler:751751751@localhost:5432/thehorned_db?sslmode=disable

migrate_down:
	sqlx migrate revert --database-url postgres://Wexler:751751751@localhost:5432/thehorned_db?sslmode=disable

build:
	if [ -f "${BINARY}" ]; then \
		rm ${BINARY}; \
		echo "Deleted ${BINARY}"; \
	fi
	echo "Building binary === "
	go build -o ${BINARY} cmd/server/*.go

run: build
	./${BINARY}

stop:
	@-pkill -SIGTERM -f "./${BINARY}"
	@echo "Stopped mf, chill out"
