# Start from the official Golang base image
FROM golang:1.22.0

# Set the Current Working Directory inside the container
WORKDIR /go_stuff

# Copy the source code into the container
COPY simple_counter /go_stuff/simple_counter

# Move to the source directory
WORKDIR /go_stuff/simple_counter

# Build the Go app
RUN go build -o simple_counter

# Command to run the executable
CMD ["./simple_counter"]
