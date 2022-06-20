FROM golang:1.17-alpine as builder

WORKDIR /app/src/crud-go
ENV GOPATH=/app
COPY . /app/src/crud-go
RUN go get -u github.com/go-sql-driver/mysql
RUN go build -o main .
CMD [ "./main" ]
