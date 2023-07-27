FROM golang:1.18.1
WORKDIR /app
COPY . ./
RUN go build -buildvcs=false
EXPOSE 12001
CMD ["./broker"]
