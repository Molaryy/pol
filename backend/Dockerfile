FROM golang:1.22.2

WORKDIR /pol-back

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o pol-back

EXPOSE ${POL_BACK_PORT}

CMD ["./pol-back"]
