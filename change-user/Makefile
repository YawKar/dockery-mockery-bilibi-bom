.PHONY: build run-from-nobody run-from-root clean

build:
	docker build -t change-user:latest .

data/flag.txt:
	sudo rm -rf data
	sudo ./generate_data.sh

run-from-nobody: build data/flag.txt
	docker run -v ./data:/data -u nobody change-user:latest

run-from-root: build data/flag.txt
	docker run -v ./data:/data -u root change-user:latest

clean:
	docker ps -a --format "{{ .ID }}" | xargs -I % docker stop % | xargs -I % docker rm %
	sudo rm -rf data
