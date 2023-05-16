FROM golang:1.20

WORKDIR ./weather

COPY . ./

RUN go build -o /weather

CMD ["/weather"]