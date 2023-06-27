# Builds the project
build:
	./build-aws-lambda.sh

install:
	go install

clean:
	go clean

.PHONY:  clean build