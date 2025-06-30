all:
	go build -o sniffer main.go

clean:
	rm -rf sniffer

run:
	./sniffer

list-interfaces:
	@echo "Available network interfaces:"
	@ls /sys/class/net/ | sed 's/^/  /'