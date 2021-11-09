FROM golang:alpine 

WORKDIR /app

COPY . .

RUN go build -o "event-project" -ldflags "-w -s"

EXPOSE 8080
ENTRYPOINT [ "./event-project" ]
