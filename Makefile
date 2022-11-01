local-build:
	@go build -v -o bin/restmock ./cmd/restmock/

start: local-build
	./bin/restmock --config ./config/conf.yaml

docker-build:
	@docker build -f ./build/Dockerfile -t restmock .

docker-start:
	@docker ps | grep restmock | awk '{print $$1}'   | xargs docker stop
	@docker ps -a | grep restmock | awk '{print $$1}'| xargs docker rm
	@docker images | grep none | awk '{print $$3}' | xargs docker rmi
	@docker run -d -p 8080:8080 restmock > /dev/null