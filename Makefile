all: build up

up: build
	docker compose -f docker-compose.yml up -d

build:
	mkdir -p ./frontend/node_modules/
	mkdir -p ./Database
	docker compose -f docker-compose.yml build

down:
	docker compose -f docker-compose.yml down -v

clean: down remove
	docker compose -f docker-compose.yml down --rmi all
	docker run -it --rm -v $(pwd):/workspace alpine sh -c "cd /workspace && rm -rf Database/"

re: down up

docker-clean:
	docker stop $$(docker ps -q)
	docker rm $$(docker ps -aq)
	docker rmi $$(docker images -aq)
	docker system prune -af