FROM golang:latest AS builder
WORKDIR /go/src/github.com/vaskoz/gophercises-phone
ADD . /go/src/github.com/vaskoz/gophercises-phone
RUN cd /go/src/github.com/vaskoz/gophercises-phone/cmd/mariadbsql \
&& CGO_ENABLED=0 GOOS=linux go build -o mariadbsql -a -tags netgo -ldflags '-w'

FROM scratch
WORKDIR /app
COPY --from=builder /go/src/github.com/vaskoz/gophercises-phone/cmd/mariadbsql/mariadbsql /app
CMD ["/app/mariadbsql"]
