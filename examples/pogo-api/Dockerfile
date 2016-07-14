FROM golang:1.7
ADD main.go .
RUN go get -v github.com/montanaflynn/is-pokemon-go-online
RUN go build -v -o ./pogo-api .
CMD ["./pogo-api"]
EXPOSE 8080
