FROM golang:1.16-alpine
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /synapsis-test

EXPOSE 5051

CMD [ "/synapsis-test" ]