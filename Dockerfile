FROM golang:1.22.0

WORKDIR /go_stuff

COPY simple_counter /go_stuff/simple_counter

WORKDIR /go_stuff/simple_counter

RUN go build -o simple_counter

CMD ["./simple_counter"]
