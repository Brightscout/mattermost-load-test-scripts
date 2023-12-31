GO ?= $(shell command -v go 2> /dev/null)

build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -gcflags "all=-N -l" -trimpath -o dist/load-tester;

create_users:
	dist/load-tester create_users

clear_store:
	dist/load-tester clear_store

create_channels:
	dist/load-tester create_channels

create_dm_and_gm:
	dist/load-tester create_dm_and_gm

create_posts:
	k6 run k6/createPosts.js

check-style: 
	@if ! [ -x "$$(command -v golangci-lint)" ]; then \
		echo "golangci-lint is not installed. Please see https://github.com/golangci/golangci-lint#install for installation instructions."; \
		exit 1; \
	fi; \

	@echo Running golangci-lint
	golangci-lint run --timeout 15m0s ./...
