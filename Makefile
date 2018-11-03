
prepare:
	go get -u github.com/golang/dep/cmd/dep

restore:
	dep ensure -v

run:
	docker-compose up -d
	@make logs

stop:
	docker-compose down

build:
	@make restore
	docker-compose build

logs:
	docker-compose logs -f

put-resources:
	curl -XPUT http://localhost:28082/resources/sidecar -d @sample/resource/sidecar.json

test:
	go test -v ./...
