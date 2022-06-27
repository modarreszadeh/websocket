FROM golang

WORKDIR /app
COPY . .
RUN go get

RUN go build -o socket.sh

EXPOSE 8080

CMD [ "./socket.sh" ]