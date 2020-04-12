APP=wait-pg-db
BIN=bin/$(APP)

.PHONY: build
build d:
	@GOOS=linux go build -a -o $(BIN)
