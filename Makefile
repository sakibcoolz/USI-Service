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
