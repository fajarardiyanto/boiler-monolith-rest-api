.PHONY:

help:
	@echo "See Makefile"
tidy:
	@go mod tidy
run :
	@ go run .
build:
	@ go build .
scan:
	@scripts/gosec.sh
docker:
	@scripts/docker.sh