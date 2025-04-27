.PHONY: build run run-interactive clean

build:
	docker build -t stdin:latest .

run: build
	docker run stdin:latest

run-interactive: build
	docker run -i stdin:latest # `-i` allows us to use stdin and pass it to container

run-tty-interactive: build
	docker run -it stdin:latest # `-i` allows us to use stdin and pass it to container

clean:
	docker ps -a --format "{{ .ID }}" | xargs -I % docker stop % | xargs -I % docker rm %
