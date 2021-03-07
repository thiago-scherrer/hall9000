FROM golang:1.16

WORKDIR /app

COPY . /app

RUN go build 

CMD ["./hall9000"]
