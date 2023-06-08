# Builds the project
build:
	./build.sh

install:
	go install

clean:
	go clean

.PHONY:  clean build