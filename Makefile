all: build

build: main.go
	go build -o azareal

win: main.go
	go build -o azareal.exe

run:
	go run main.go

clean:
	rm azareal*
