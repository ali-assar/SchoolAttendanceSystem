BIN_DIR = ./bin

EXECUTABLE = attendence_api
EXECUTABLE_reduce = attendence_api_reduce
RPI_EXECUTABLE = attendence_raspberry_api

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

sqlc:
	@echo "Generating database code ... "
	@ sqlc/sqlc generate

build-reduce: $(BIN_DIR)
	@echo "Building the main application..."
	@/usr/local/go/bin/go build -ldflags="-s -w" -o $(BIN_DIR)/$(EXECUTABLE_reduce) 

build: $(BIN_DIR)
	@echo "Building the main application..."
	@/usr/local/go/bin/go build -o $(BIN_DIR)/$(EXECUTABLE) 

run: build
	@echo "Running the main application..."
	@./$(BIN_DIR)/$(EXECUTABLE)

build-raspberry: $(BIN_DIR)
	@echo "Cross-compiling for Raspberry Pi..."
	@GOOS=linux GOARCH=arm GOARM=7 go build -o $(BIN_DIR)/$(RPI_EXECUTABLE) -v

seed:
	@echo "seeding the database..."
	@go run seed/seed.go

.PHONY: seed
