APP=todo-backend
BIN=/bin

swagger:
	swag init

build:
	go build -o .$(BIN)/$(APP) main.go

bin:
	.$(BIN)/$(APP)

clean:
ifeq ($(OS),Windows_NT)
	rm $(BIN)
else
	rm -rf $(BIN)
endif

run:
	@echo Building the swagger documentation...
	make swagger
	@echo Building the go API from the source code...
	make build
	@echo Running the API
	.$(BIN)/$(APP)