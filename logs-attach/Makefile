.PHONY: build logs attach clean

build:
	docker build -t logs-attach:latest .

.cached_id: build
	docker run -d logs-attach:latest > .cached_id

logs: .cached_id
	sleep 5 && docker logs `cat .cached_id`

attach: .cached_id
	docker attach `cat .cached_id`

clean:
	docker ps -a --format "{{ .ID }}" | xargs -I % docker stop % | xargs -I % docker rm %
