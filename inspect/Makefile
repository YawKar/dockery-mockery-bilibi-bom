.PHONY: build run run-and-inspect inspect clean

build:
	docker build -t inspect:latest .

run: build
	docker run -d inspect:latest sleep 100 > .cached-id

inspect: run
	docker inspect `cat .cached-id`

run-and-inspect: build
	docker inspect `docker run -d inspect:latest`

clean:
	docker ps -a --format "{{ .ID }}" | xargs -I % docker stop % | xargs -I % docker rm %
	rm -f .cached-id
