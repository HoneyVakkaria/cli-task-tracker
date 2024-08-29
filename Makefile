TARGET = task-tracker

build:
	go build -o $(TARGET)

clean:
	rm -f $(TARGET)

install: build
	cp $(TARGET) /usr/local/bin/
	rm -f $(TARGET)

uninstall:
	rm -f /usr/local/bin/$(TARGET)

.PHONY: all build install clean uninstall
all: build
