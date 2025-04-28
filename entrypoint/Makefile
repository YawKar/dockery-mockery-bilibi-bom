.PHONY: build run run-with-arg
build:
	(docker build -t entrypoint:try . && echo "Built successful") || echo "Build failed"

run: build
	docker run entrypoint:try

run-with-arg: build
	docker run entrypoint:try YawKar

run-with-override: build
	docker run --entrypoint echo entrypoint:try echoed
