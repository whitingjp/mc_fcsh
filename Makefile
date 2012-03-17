all: bindir bin/fcsh_wrap.exe bin/fcsh_server.exe
clean:
	rm -rf bin
bindir:
	mkdir -p bin
bin/fcsh_wrap.exe: client.go
	go build -o $@ $^
bin/fcsh_server.exe: server.go
	go build -o $@ $^
format:
	gofmt -w client.go
	gofmt -w server.go