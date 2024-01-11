FROM golang:1.19.1
WORKDIR /app
COPY . ./
RUN go build -buildvcs=false
EXPOSE 12001
CMD ["./broker", "0.0.0.0:12001"]
