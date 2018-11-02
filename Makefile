
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
