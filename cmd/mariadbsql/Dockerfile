FROM golang:latest AS builder
ADD . /go/src/github.com/vaskoz/gophercises-phone
WORKDIR /go/src/github.com/vaskoz/gophercises-phone/cmd/mariadbsql
RUN CGO_ENABLED=0 GOOS=linux go build -o mariadbsql -a -tags netgo -ldflags '-w'

FROM scratch
WORKDIR /app
COPY --from=builder /go/src/github.com/vaskoz/gophercises-phone/cmd/mariadbsql/mariadbsql .
CMD ["./mariadbsql"]
