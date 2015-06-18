all:
	@echo "Usage: $ make [build|clean]"

build:
	@./build.sh

clean:
	rm bin/*

run:
	./bin/Reporter