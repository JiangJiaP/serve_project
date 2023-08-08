# 使用官方的 Golang 镜像作为基础镜像
FROM golang:latest as builder

# 设置工作目录
WORKDIR /app

# 将项目中的 go.mod 和 go.sum 文件复制到工作目录
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 将项目中的源代码复制到工作目录
COPY . .

# 编译 Go 程序
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# 使用一个较小的基础镜像，以减小最终镜像的体积
FROM alpine:latest

# 将编译好的二进制文件从 builder 镜像复制到当前镜像
COPY --from=builder /app/myapp /app/

# 设置工作目录
WORKDIR /app

# 暴露端口（根据你的应用程序需要暴露的端口进行修改）
EXPOSE 8080

# 运行 Go 程序
CMD ["./myapp"]