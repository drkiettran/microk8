FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o hello-web .
RUN ls /app
CMD ["/app/hello-web"]
