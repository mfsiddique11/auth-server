FROM golang:1.13
RUN mkdir /app
Add . /app
WORKDIR /app
RUN go build -o main .
EXPOSE 5000
CMD ["/app/main"]
