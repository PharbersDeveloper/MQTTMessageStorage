#源镜像
FROM golang:1.12.4-alpine

#作者
MAINTAINER Pharbers "pqian@pharbers.com"

# 安装系统级依赖
RUN apk add --no-cache git gcc musl-dev mercurial bash gcc g++ make pkgconfig openssl-dev

# 设置工程配置文件的环境变量
ENV PKG_CONFIG_PATH /usr/lib/pkgconfig
ENV MQTT_HOME $GOPATH/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources
ENV BM_KAFKA_CONF_HOME $GOPATH/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources/resource/kafkaconfig.json
ENV GO111MODULE on
ENV LOGGER_USER "Alex"
ENV LOGGER_DEBUG "true"
ENV LOG_PATH /mqtt.log

# 以LABEL行的变动(version的变动)来划分(变动以上)使用cache和(变动以下)不使用cache
LABEL NtmPods.version="0.0.3" maintainer="Alex"

# 下载kafka
RUN git clone https://github.com/edenhill/librdkafka.git $GOPATH/librdkafka

WORKDIR $GOPATH/librdkafka
RUN ./configure --prefix /usr  && \
make && \
make install

# 下载依赖

RUN git clone https://github.com/PharbersDeveloper/MQTTMessageStorage.git  $GOPATH/src/github.com/PharbersDeveloper/MQTTMessageStorage

# 构建可执行文件
RUN cd $GOPATH/src/github.com/PharbersDeveloper/MQTTMessageStorage && \
    go build && go install

# 暴露端口
EXPOSE 6542

# 设置工作目录
WORKDIR $GOPATH/bin

ENTRYPOINT ["MQTTStorage"]
