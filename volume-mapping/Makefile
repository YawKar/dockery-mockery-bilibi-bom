.PHONY: build run run-using-mount clean

build:
	docker build -t volume-mapping:latest .

run: build
	docker run -v ./data:/data volume-mapping:latest /data

run-using-mount: build
	docker run --mount type=bind,source=./data,target=/data volume-mapping:latest /data

run-using-volume: build create-volume
	docker run --mount type=volume,source=test_volume_aeaazaoenuh,target=/data volume-mapping:latest /data

create-volume:
	docker volume create test_volume_aeaazaoenuh

clean:
	docker ps -a --format "{{ .ID }}" | xargs -I % docker stop % | xargs -I % docker rm %
	docker volume rm test_volume_aeaazaoenuh
