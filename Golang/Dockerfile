FROM golang:1.23

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ../source .

RUN go mod tidy

#RUN go build -o bin .

#ENTERPOINT [ "/app/bin" ]