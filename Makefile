GO ?= $(shell command -v go 2> /dev/null)

build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -gcflags "all=-N -l" -trimpath -o dist/plugin-build;

create_users:
	dist/plugin-build create_users

clear_store:
	dist/plugin-build clear_store
