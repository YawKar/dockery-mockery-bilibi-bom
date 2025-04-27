.PHONY: build run clean

build:
	docker build -t port-mapping:latest .

run: build
	docker run -d -p 80:5050 port-mapping:latest # 80 on host machine to 5050 inside container

ping:
	(curl localhost:80/ping && printf "\nSuccess!\n") || printf "\nFailure!\n"

clean:
	docker ps -a --format "{{ .ID }}" | xargs -I % docker stop % | xargs -I % docker rm %
