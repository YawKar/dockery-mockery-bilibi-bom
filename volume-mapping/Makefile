.PHONY: build run clean

build:
	docker build -t volume-mapping:latest .

run: build
	docker run -v ./data:/data volume-mapping:latest /data

clean:
	docker ps -a --format "{{ .ID }}" | xargs -I % docker stop % | xargs -I % docker rm %
