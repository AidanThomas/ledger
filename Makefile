cmd ?= tui

start:
	@go run cmd/${cmd}/main.go
.PHONY: start
