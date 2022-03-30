install:
	go install github.com/swaggo/swag/cmd/swag@latest

generate:
	swag init -ot go --parseDependency

docker-clean:
	docker system prune -a

docker-build:
	docker buildx build --platform linux/amd64 -t kedai-api-gateway .
	docker tag kedai-api-gateway alganbr/kedai-api-gateway
	docker push alganbr/kedai-api-gateway

docker-rebuild:
	make docker-clean
	make docker-build