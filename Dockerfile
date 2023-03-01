FROM golang:1.19-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd /app/cmd
COPY ./db /app/db

RUN go build ./cmd/main.go 
EXPOSE 2001

CMD [ "./main" ]