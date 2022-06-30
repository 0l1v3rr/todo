FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app

# download the dependencies
RUN go mod download

# port
EXPOSE 8080

# build
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN make swagger
RUN make build

# run
CMD ["./bin/todo-backend"]