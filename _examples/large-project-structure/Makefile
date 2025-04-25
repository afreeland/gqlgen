BINARY_NAME=setup
SOURCE_FILE=setup.go

all: build run clean

build:
	go build -o $(BINARY_NAME) $(SOURCE_FILE)

run:
	./$(BINARY_NAME) --destdir=$(DESTDIR)

clean:
	rm -f $(BINARY_NAME)
