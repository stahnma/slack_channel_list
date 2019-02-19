

fmt:
	go fmt slack_channel_list.go

build:
	go build slack_channel_list.go

install:
	sudo install -p -m755 slack_channel_list /usr/local/bin
