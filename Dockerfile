FROM golang:1.16-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
EXPOSE 3000
CMD ["/app/main"]