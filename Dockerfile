FROM golang:1.23 AS builder

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GIN_MODE=release \
  GOPROXY=https://goproxy.cn,direct

WORKDIR /app
ADD go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o wechat-robot-admin-backend


FROM alpine:latest

ENV GIN_MODE=release \
  TZ=Asia/Shanghai

# 安装 Docker CLI 和其他必要工具
RUN apk add --no-cache docker-cli ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/wechat-robot-admin-backend ./
COPY --from=builder /app/robot.sql ./robot.sql

EXPOSE 9000

CMD ["/app/wechat-robot-admin-backend"]