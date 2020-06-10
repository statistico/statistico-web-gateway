# Step 1
FROM golang:1.12.9-alpine AS builder
RUN apk update && apk add git gcc g++
ENV GO111MODULE=on
WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/rest

# Step 2
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /opt
COPY ./api ./opt/api
COPY ./bin ./bin
COPY --from=builder /app/rest .

CMD ["/bin/sh"]