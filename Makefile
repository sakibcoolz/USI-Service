include compose/.env
include variables.mk

run:
	DB_HOST=$(localhost) \
	DB_PORT=$(dbport) \
	DB_USERNAME=$(dbuser) \
	DB_PASSWORD=$(dbpassword) \
	DB_NAME=$(database) \
	DB_SCHEME=$(database) \
	SERVICE_HOST=$(localhost) \
	SERVICE_PORT=$(port) \
	$(GORUN) cmd/main.go

docker-image:
	make -C docker docker-image

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(build) ./cmd/

docker-start:
	make -C compose start

docker-stop:
	make -C compose stop

kick-start: docker-image docker-start

clean:
	@docker rmi -f $(service)
	@docker rmi -f $(database)
	@docker image prune -f
	

.PHONY: run