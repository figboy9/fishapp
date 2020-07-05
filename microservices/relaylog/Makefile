
PJT_NAME = $(notdir $(PWD))
NET = fishapp-net
SCHEMA_PATH = $$(dirname $$(dirname $$(pwd)))/schema

SVC = relaylog

createnet:
	docker network create $(NET)

proto:
	docker run --rm --name protoc -v $(PWD)/pb:/pb -v $(SCHEMA_PATH):/proto ezio1119/protoc \
	-I/proto \
	-I/go/src/github.com/envoyproxy/protoc-gen-validate \
	--go_out=plugins=grpc:/pb \
	--validate_out="lang=go:/pb" \
	event.proto chat.proto post.proto

test:
	docker-compose exec $(SVC) sh -c "go test -v -coverprofile=cover.out ./... && \
	go tool cover -html=cover.out -o ./cover.html" && \
	open ./src/cover.html

up:
	docker-compose up -d

build:
	docker-compose build

down:
	docker-compose down

exec:
	docker-compose exec $(SVC) sh

logs:
	docker logs -f --tail 100 $(PJT_NAME)_$(SVC)_1

natslogs:
	docker logs -f --tail 100 $(PJT_NAME)_$(NATS_SVC)_1

rmvol:
	docker-compose down -v