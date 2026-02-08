# 阶段一：构建阶段
FROM golang:1.25-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制依赖文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 编译二进制文件，禁用 CGO
RUN CGO_ENABLED=0 go build -o /app/server ./cmd/server/main.go

# 阶段二：运行阶段
FROM alpine:latest

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/server /app/server

# 暴露服务端口
EXPOSE 8080

# 设置入口命令
ENTRYPOINT ["/app/server"]
