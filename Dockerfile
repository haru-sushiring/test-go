## ベースイメージ
FROM golang:1.23

# 作業ディレクトリを設定
WORKDIR /app

# 必要なファイルをコピー
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ビルド
# RUN go build -o main .
# RUN go build -o main . && ls -l main
# RUN go build -o main . && chmod +x main && ls -l main
# RUN go build -o /app/main . && chmod +x /app/main && ls -l /app
RUN go build -o /app/main .

# アプリケーションを実行
# CMD ["./main"]
CMD ["/app/main"]