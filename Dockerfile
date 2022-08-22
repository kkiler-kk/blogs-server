# syntax = docker/dockerfile:experimental
# 编译
FROM golang:1.16-alpine3.13
COPY . /build/
WORKDIR /build
RUN --mount=type=cache,target=/go/pkg,id=go_pkg,sharing=locked \
    GOPROXY=https://goproxy.cn GOSUMDB=off CGO_ENABLED=1 GOOS=linux GOARCH=amd64  go build -ldflags '-w -s' -o server
# 运行阶段 基础镜像加上时区
FROM alpine:3.13
# 设置源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
# 设置时区
    apk add --no-cache tzdata  && echo "Asia/Shanghai" > /etc/timezone && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# 从编译阶段的中拷贝编译结果到当前镜像中
COPY --from=0 /build/server /
# 拷贝配置文件
COPY --from=0 /build/config/app.ini /data/config/app.ini
COPY --from=0 /build/config/app.test.ini /data/config/app.test.ini
COPY --from=0 /build/config/acs.model.conf /data/config/acs.model.conf
WORKDIR /data
ENTRYPOINT ["/server"]