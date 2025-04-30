.PHONY: run-registry build push-image make-tag pull-image clean

run-registry:
	docker run -d -p 5000:5000 --name registry registry:2

build:
	docker build -t image:latest .

make-tag:
	docker image tag image:latest localhost:5000/image:latest # like 'docker.io/image:latest'

push-image: run-registry build make-tag
	docker push localhost:5000/image:latest

pull-image: run-registry build make-tag push-image
	docker pull localhost:5000/image:latest

clean:
	docker ps -a --format "{{ .ID }}" | xargs -I % docker stop % | xargs -I % docker rm %
	docker images --format "{{ .Repository }}" | grep "localhost" | xargs -I % docker image rm %
