FROM golang:1.19 as builder

WORKDIR /go/src/app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENALBED=0 GOOS=linux GOARCH=amd64 go build -v -o server ./cmd

WORKDIR /go/src/app

FROM gcr.io/distroless/static:nonroot as runner

COPY --from=builder /go/src/app/server /server

ENTRYPOINT ["/server"]
