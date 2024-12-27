.PHONY: clean ps umount shell build test

build:
	go build

ps:
	sudo mount --make-rprivate /
	sudo ./ccrun run /bin/busybox ps

shell:
	sudo mount --make-rprivate /
	sudo ./ccrun run /bin/busybox sh

clean:
	go clean

test:
	sudo go test ./...