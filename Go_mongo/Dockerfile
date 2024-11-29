# 使用 Go 1.23 的官方映像
FROM golang:1.23

# 設定工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum，然後下載依賴
COPY ./source ./

RUN go mod init "example"
RUN go mod tidy

# 編譯 Go 應用程式
RUN go build -o main .

# 設定執行命令
CMD ["./main"]
