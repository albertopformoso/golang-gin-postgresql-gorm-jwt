create-db:
	docker-compose -f ./db/docker-compose.yml up -d --build

remove-db:
	docker-compose -f ./db/docker-compose.yml down

prune-db:
	docker-compose -f ./db/docker-compose.yml down
	docker volume prune --force

halt-db:
	make remove-db
	make create-db

restart-db:
	make prune-db
	make create-db

build:
	docker-compose -f ./docker/docker-compose.yml up -d --build

delete:
	docker-compose -f ./docker/docker-compose.yml down

test:
	docker-compose -f ./docker/docker-compose-test.yml up -d --build
	sleep 5s
	docker logs go-api-test
	docker rm -f go-api-test
	docker rmi -f go-alpine-test:1.0.0