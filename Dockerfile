# 构建golang环境的基础镜像
FROM registry.cn-hangzhou.aliyuncs.com/xzjs/golang:1.14 as build

# 容器环境变量添加，会覆盖默认的变量值
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作区
WORKDIR /go/forum

# 把全部文件添加到/go/forum目录
ADD . .

# 编译：把main.go编译成可执行的二进制文件，命名为app
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o ./app

# 运行：使用scratch作为基础镜像
FROM registry.cn-hangzhou.aliyuncs.com/xzjs/alpine:3.12.3 as prod

RUN apk --no-cache add ca-certificates && rm -rf /var/cache/apk/* /tmp/*

# 在build阶段复制可执行的go二进制文件app
COPY --from=build /go/forum/app /
EXPOSE 8888

# 启动服务
CMD ["/app"]