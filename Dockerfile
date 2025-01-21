## ベースイメージ
FROM golang:1.20

# 作業ディレクトリを設定
WORKDIR /app

# 必要なファイルをコピー
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ビルド
RUN go build -o main .

# アプリケーションを実行
CMD ["./main"]