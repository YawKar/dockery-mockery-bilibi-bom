.PHONY: build run run-with-exec clean

build:
	docker build -t commands:latest .

run: build
	docker run commands:latest

run-with-exec: build
	docker exec `docker run -d commands:latest sleep 1` cat /etc/os-release

clean:
	docker ps -a --format "{{ .ID }}" | xargs -I % docker stop % | xargs -I % docker rm %
