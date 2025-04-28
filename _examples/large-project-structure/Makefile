BINARY_NAME=setup
SOURCE_FILE=setup.go

all: build run clean

build:
	go build -o $(BINARY_NAME) $(SOURCE_FILE)

run:
	./$(BINARY_NAME) --destdir=$(DESTDIR)
	cd ./main
	go generate ./...

clean:
	rm -f $(BINARY_NAME)
