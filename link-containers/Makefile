.PHONY: build run-composition clean run-server run-client

build:
	docker build -f server.Dockerfile -t server:latest .
	docker build -f client.Dockerfile -t client:latest .

run-server: build
	docker run -d --name server server:latest 8080

run-client: build
	docker run -d --name client --link server:server_host client:latest server_host 8080

run-composition: run-server run-client
	docker logs client
	echo "What Docker has added to /etc/hosts of client container:"
	docker cp client:/etc/hosts -

clean:
	docker ps -a --format "{{ .ID }}" | xargs -I % docker stop % | xargs -I % docker rm %
