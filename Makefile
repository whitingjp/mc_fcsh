all: bindir bin/mc_fcsh bin/dj_fcsh
clean:
	rm -rf bin
bindir:
	mkdir -p bin
bin/mc_fcsh: mc_fcsh.go
	go build -o $@ $^
bin/dj_fcsh: dj_fcsh.go
	go build -o $@ $^
format:
	gofmt -w mc_fcsh.go
	gofmt -w dj_fcsh.go