PJT_NAME = $(notdir $(PWD))
NET = fishapp-net
# makefileのシェル関数で実行するとPermission deniedになるので、bashで直接実行
SCHEMA_PATH = $$(dirname $$(dirname $$(pwd)))/schema
SVC = api-gateway

proto:
	docker run --rm --name protoc -v $(PWD)/pb:/pb -v $(SCHEMA_PATH):/proto ezio1119/protoc \
	-I/proto \
	-I/go/src/github.com/envoyproxy/protoc-gen-validate \
	--go_out=plugins=grpc:/pb \
	--validate_out="lang=go:/pb" \
	chat.proto post.proto user.proto image.proto event.proto

gql:
	docker run --rm --name gqlgen -w /app -v $(PWD):/app -v $(SCHEMA_PATH)/api-gateway:/app/schema ezio1119/gqlgen \
	go run github.com/99designs/gqlgen

test:
	docker-compose exec $(SVC) sh -c "go test -v -coverprofile=cover.out ./... && \
	go tool cover -html=cover.out -o ./cover.html" && \
	open ./src/cover.html

up:
	docker-compose up -d $(SVC)

build:
	docker-compose build

down:
	docker-compose down

exec:
	docker-compose exec $(SVC) sh

logs:
	docker logs -f --tail 100 $(PJT_NAME)_$(SVC)_1
	