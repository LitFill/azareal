COMPILER ?= go
BINNAME ?= azareal

BUILDCMD ?= $(COMPILER) build
OUTPUT ?= -o $(BINNAME)
FLAGS ?= -v

RUNCMD ?= $(COMPILER) run

all: build

build: main.go
	$(BUILDCMD) $(OUTPUT) $(FLAGS)

win: main.go
	$(BUILDCMD) $(OUTPUT).exe $(FLAGS)

run: main.go
	$(RUNCMD) $(FLAGS) $^

clean:
	rm $(BINNAME)*
