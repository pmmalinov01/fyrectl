help:
	@echo "Print Help"

compileMac:
	go build -o fyrectl main.go

compileLinux:
	GOOS=linux GOARCH=386 go build -o fyrectl main.go

list:
	./fyrectl list
test: compileMac list
