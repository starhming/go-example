FROM golang:1.19 as builder
# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

LABEL maintainer="minghai1024"

WORKDIR /build

# 将代码复制到容器中
RUN go mod init hm1024/http-server
COPY . .
RUN go mod tidy

# 将我们的代码编译成二进制可执行文件app
RUN go build -o http-server .

# use a minimal alpine image
FROM alpine:3.7
WORKDIR /app
# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/http-server /app
EXPOSE 8080
ENTRYPOINT ["/app/http-server"]