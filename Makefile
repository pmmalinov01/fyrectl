help:
	@echo "Print Help"

compileMac:
	go build -o fyrectl main.go

compileLinux:
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go

list:
	./fyrectl list
test: compileMac list
