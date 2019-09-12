# builder 源镜像
FROM golang:1.12.4-alpine as builder

# 作者
MAINTAINER Pharbers "pqian@pharbers.com"

# 安装系统级依赖
RUN echo http://mirrors.aliyun.com/alpine/edge/main > /etc/apk/repositories \
&& echo http://mirrors.aliyun.com/alpine/edge/community >> /etc/apk/repositories \
&& apk update \
&& apk add --no-cache bash git gcc g++ openssl-dev librdkafka-dev pkgconf

# 环境变量
#ENV PKG_CONFIG_PATH /usr/lib/pkgconfig
ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

# 下载项目镜像
#RUN git clone https://github.com/PharbersDeveloper/MQTTMessageStorage.git  $GOPATH/src/github.com/PharbersDeveloper/MQTTMessageStorage

# 工作目录
WORKDIR $GOPATH/src/github.com/PharbersDeveloper/MQTTMessageStorage
COPY . .

# go build 编译项目
RUN go build




# prod 源镜像
FROM alpine:latest

# 作者
MAINTAINER Pharbers "pqian@pharbers.com"

# 安装 主要 依赖
RUN echo http://mirrors.aliyun.com/alpine/edge/main > /etc/apk/repositories \
&& echo http://mirrors.aliyun.com/alpine/edge/community >> /etc/apk/repositories \
&& apk update \
&& apk add --no-cache bash git gcc g++ openssl-dev librdkafka-dev pkgconf

# 环境变量
#ENV PKG_CONFIG_PATH /usr/lib/pkgconfig
ENV MQTT_HOME /go/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources
ENV BM_KAFKA_CONF_HOME /go/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources/resource/kafkaconfig.json
ENV PROJECT_NAME MQTTMessageStorage
ENV BP_LOG_TIME_FORMAT "2006-01-02 15:04:05"
ENV BP_LOG_OUTPUT /go/log/mqtt.log
ENV BP_LOG_LEVEL info

WORKDIR /go/log

WORKDIR /go/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources/resource

# 提取资源文件
COPY --from=0 /go/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources/resource/kafkaconfig.json ./
COPY --from=0 /go/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources/resource/redisconfig.json ./
COPY --from=0 /go/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources/resource/routerconfig.json ./
COPY --from=0 /go/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources/resource/service-def.yaml ./

WORKDIR /go/bin

# 提取执行文件
COPY --from=0 /go/src/github.com/PharbersDeveloper/MQTTMessageStorage/MQTTStorage ./

EXPOSE 6542
ENTRYPOINT ["./MQTTStorage"]
