FROM golang:alpine as builder

# 设置go mod proxy 国内代理
# 设置golang path
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=1
WORKDIR /dist
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
COPY . .
RUN go env && go list && go build -o app main.go

EXPOSE 8002
# 启动容器时运行的命令
CMD ["/dist/app env=prod"]

