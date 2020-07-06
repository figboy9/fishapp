PJT_NAME = $(notdir $(PWD))
SCHEMA_PATH = $$(dirname $$(dirname $$(pwd)))/schema
SVC = image
NET = fishapp-net
DB_SVC = image
DB_NAME = image_DB
GRPC_SVC = ImageService

proto:
	docker run --rm --name protoc -v $(PWD)/pb:/pb -v $(SCHEMA_PATH):/proto ezio1119/protoc \
	-I/proto \
	-I/go/src/github.com/envoyproxy/protoc-gen-validate \
	--go_out=plugins=grpc:/pb \
	--validate_out="lang=go:/pb" \
	image.proto

newsql:
	docker run --rm --name newsql -v $(PWD)/db/sql:/sql \
	migrate/migrate:latest create -ext sql -dir ./sql ${a}

test:
	docker-compose exec $(SVC) sh -c "go test -v -coverprofile=cover.out ./... && \
	go tool cover -html=cover.out -o ./cover.html" && \
	open ./src/cover.html

cli:
	docker run --rm --name grpc_cli --net $(NET) znly/grpc_cli \
	call $(SVC):50051 $(SVC).$(GRPC_SVC).$(m) "$(q)"

migrate:
	docker run --rm --name migrate --net $(NET) \
	-v $(PWD)/db/sql:/sql migrate/migrate:latest \
	-path /sql/ -database "mysql://root:password@tcp($(DB_SVC):3306)/$(DB_NAME)" $(MIGRATE)

waitdb: updb
	docker run --rm --name dockerize --net $(NET) jwilder/dockerize \
	-timeout 60s \
	-wait tcp://$(DB_SVC):3306

migrate: waitdb
	docker run --rm --name migrate --net $(NET) \
	-v $(PWD)/db/sql:/sql migrate/migrate:latest \
	-path /sql/ -database "mysql://root:password@tcp($(DB_SVC):3306)/$(DB_NAME)" $(MIGRATE)

sqldoc: migrate
	docker run --rm --name tbls --net $(NET) -v $(PWD)/db:/work ezio1119/tbls \
	doc -f -t svg mysql://root:password@$(DB_SVC):3306/$(DB_NAME) ./

up:
	docker-compose up -d

updb:
	docker-compose up -d $(DB_SVC)

build:
	docker-compose build

down:
	docker-compose down

exec:
	docker-compose exec $(SVC) sh

logs:
	docker logs -f --tail 100 $(PJT_NAME)_$(SVC)_1

rmvol:
	docker-compose down -v
