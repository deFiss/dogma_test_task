FROM golang:1.16
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o api ./cmd/main.go
CMD [ "./api" ]