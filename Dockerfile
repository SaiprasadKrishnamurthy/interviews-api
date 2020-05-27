FROM golang:latest
LABEL maintainer="Sai Kris"

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod .
COPY go.sum .

# Add this go mod download command to pull in any dependencies
RUN go mod download

# Our project will now successfully build with the necessary go libraries included.
RUN go build -o main .

# Our start command which kicks off
# our newly created binary executable
CMD ["/app/main"]
