FROM golang:1.10 as builder

## Create a directory and Add Code
RUN mkdir -p /go/src/github.com/orvice/monitor-client
WORKDIR /go/src/github.com/orvice/monitor-client
ADD .  /go/src/github.com/orvice/monitor-client

RUN go get
RUN CGO_ENABLED=0 go build


FROM orvice/go-runtime

COPY --from=builder /go/src/github.com/orvice/monitor-client/monitor-client .

ENTRYPOINT [ "./monitor-client" ]