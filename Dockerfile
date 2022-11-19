FROM golang:1.19-apline AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main main.go

EXPOSE 80

ENTRYPOINT [ "./main" ]


