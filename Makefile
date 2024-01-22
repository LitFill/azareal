COMPILER := go
BINNAME := azareal

BUILDCMD := $(COMPILER) build
OUTPUT := -o $(BINNAME)
FLAGS := -v

RUNCMD := $(COMPILER) run

.PHONY: all build run clean win help

all: build ## Build the binary for linux

build: main.go ## Actually build the binary
	$(BUILDCMD) $(OUTPUT) $(FLAGS)

win: main.go ## Build the binary for windows
	$(BUILDCMD) $(OUTPUT).exe $(FLAGS)

run: main.go ## Run the main.go
	$(RUNCMD) $(FLAGS) $^

clean: ## Clean up
	rm -f $(BINNAME)*

help: ## Prints help for targets with comments
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
