.PHONY: up
up: 
	docker-compose up -d

.PHONY: down
down:
	docker-compose down -v

.PHONY: localstack
localstack:
	docker-compose exec -it localstack bash

.PHONY: s3
s3:
	docker-compose exec localstack awslocal s3 mb s3://sample-bucket