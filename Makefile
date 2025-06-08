

# App settings
APP_NAME     := taskman
MAIN_SRC     := ./cmd/taskman/main.go

# OS-specific configs
ifeq ($(OS),Windows_NT)
	EXEC_EXT := .exe
	RM := cmd /c "del /f /q"
else
	EXEC_EXT :=
	RM := rm -f
endif

# Final binary name
BINARY := $(APP_NAME)$(EXEC_EXT)

# Default target
all: build

# Build the binary
build:
	go build -o $(BINARY) $(MAIN_SRC)

# Run the app
run:
	go run $(MAIN_SRC)

# Run tests
test:
	go test ./...

# Clean the build artifacts
clean:
	$(RM) $(BINARY)

dev: clean build

hello:
	@echo "Hello from $(APP_NAME)!"
