# 用于编译的镜像
FROM golang:1.16.7 as builder

# 设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 时区
ENV TZ Asia/Shanghai
RUN ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone

WORKDIR /workspace

# 拷贝代码
COPY . .

# 编译
RUN export GOPROXY=https://goproxy.io && go build -o demo

# 用于执行的镜像
FROM scratch as runner

# 支持中文
ENV LANF C.UTF-8

WORKDIR /workspace

# 从构建阶段把文件拷贝过来
# 从构建阶段的/etc/timezone和/etc/localtime文件拷贝到执行阶段的/etc/文件夹下
COPY --from=builder /etc/timezone /etc/localtime /etc/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /workspace/config.yaml /workspace/demo ./

# 声明需要开放的端口
EXPOSE 8080

# 执行
ENTRYPOINT ["./demo"]