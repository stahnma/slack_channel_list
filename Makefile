

fmt:
	go fmt slack_channel_list.go

build: fmt
	go build slack_channel_list.go

install: build
	sudo install -p -m755 slack_channel_list /usr/local/bin
