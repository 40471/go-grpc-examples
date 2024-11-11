FROM golang:1.23 as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /app/notifications ./server/main.go

FROM golang:1.23 

WORKDIR /root/
COPY --from=builder /app/notifications .
COPY --from=builder /app/proto ./proto

EXPOSE 50051

ENV RABBITMQ_HOST=rabbitmq

CMD ["./notifications"]