.PHONY: build run run-red run-green clean

build:
	docker build -t environment-variables:latest .

run: build
	docker run environment-variables:latest

run-red: build
	docker run -e COLOR=Red -e NAME=Ivan environment-variables:latest

run-green: build
	docker run -e COLOR=Green -e NAME=YawKar environment-variables:latest

clean:
	docker ps -a --format "{{ .ID }}" | xargs -I % docker stop % | xargs -I % docker rm %
