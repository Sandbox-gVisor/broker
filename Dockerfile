FROM golang:1.18.1
WORKDIR /app
COPY . ./
RUN go build
EXPOSE 12001
CMD ["./broker"]
