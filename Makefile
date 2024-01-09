all: build

build: main.go
	go build -o azareal

run:
	go run main.go
