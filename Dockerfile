FROM golang:1.20-alpine
LABEL authors="nitish"
RUN mkdir "/app"
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]