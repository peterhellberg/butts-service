.PHONY: default
default: butts

.PHONY: run
run:
	go run main.go

.PHONY: req
req:
	nats req svc.butts --count 1 ""

.PHONY: butts
butts:
	rm -rf butts butts.zip
	mkdir butts
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags -s -tags netgo,osusergo -o butts/butts-service.mac
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags -s -tags netgo,osusergo -o butts/butts-service.exe
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags -s -tags netgo,osusergo -o butts/butts-service.elf
	zip butts.zip butts/*
