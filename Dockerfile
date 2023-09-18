# Use the official Go image as the base image
FROM golang:1.20

# Set the working directory in the container
WORKDIR /app

# Copy the application files into the working directory
COPY . /app

# Build the application
#RUN protoc -Iproto --go_opt=module=github.com/olegmoney --go_out=./proto --go-grpc_opt=module=github.com/olegmoney --go-grpc_out=./proto proto/*.proto

RUN go mod tidy
RUN make proto
RUN go build -o main ./server

# Expose port 8080
EXPOSE 8080

# Define the entry point for the container
CMD ["./main"]