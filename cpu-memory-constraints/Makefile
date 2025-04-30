.PHONY: build run-cpu-30 run-half-cpu-30 run-memory-100-fail run-memory-100-pass clean

build-cpu:
	docker build -f cpu.Dockerfile -t cpu:latest .

build-memory:
	docker build -f memory.Dockerfile -t memory:latest .

run-full-cpu-30: build-cpu
	docker run --cpus `nproc` cpu:latest 30

run-half-cpu-30: build-cpu
	docker run --cpus $$((`nproc` / 2)) cpu:latest 30

run-memory-100-fail: build-memory
	docker run --memory 100M --memory-swap 100M memory:latest 99

run-memory-100-pass: build-memory
	docker run --memory 100M --memory-swap 100M memory:latest 98

clean:
	docker ps -a --format "{{ .ID }}" | xargs -I % docker stop % | xargs -I % docker rm %
