## builder image
FROM golang:1.11.6 as builder

ARG SSH_KEY
ENV GO111MODULE=on

WORKDIR /gurume

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


## final image
FROM alpine:3.9

RUN apk --no-cache add ca-certificates

COPY --from=builder gurume .
ENTRYPOINT ["./gurume"]