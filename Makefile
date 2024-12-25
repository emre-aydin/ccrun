.PHONY: clean ps umount shell build

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