FROM golang:1.11 as builder

WORKDIR /home/app
COPY go.mod go.sum ./

RUN go mod download

RUN CGO_ENABLED=0 go build -o monitor-client


FROM orvice/go-runtime:lite

COPY --from=builder /home/app/monitor-client .

ENTRYPOINT [ "./monitor-client" ]