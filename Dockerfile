# builder 源镜像
FROM golang:1.12.4-alpine as builder

# 作者
MAINTAINER Pharbers "pqian@pharbers.com"

# 安装系统级依赖
RUN apk add --no-cache bash git

# 安装 rdkafka 依赖
RUN apk add --no-cache gcc g++ make pkgconfig openssl-dev \
&& git clone https://github.com/edenhill/librdkafka.git -b v1.1.0 $GOPATH/librdkafka \
&& cd $GOPATH/librdkafka/ \
&& ./configure --prefix /usr \
&& make \
&& make install

ENV GOPROXY https://goproxy.io

WORKDIR $GOPATH

# 下载项目镜像
RUN git clone https://github.com/PharbersDeveloper/MQTTMessageStorage.git  $GOPATH/src/github.com/PharbersDeveloper/MQTTMessageStorage

# go build 编译项目
RUN go build


# prod 源镜像
FROM alpine:latest as prod

# 作者
MAINTAINER Pharbers "pqian@pharbers.com"

# 安装 主要 依赖
RUN apk --no-cache add bash git

# 安装 rdkafka 依赖
RUN apk add --no-cache gcc g++ make pkgconfig openssl-dev \
&& git clone https://github.com/edenhill/librdkafka.git -b v1.1.0 /tmp/librdkafka \
&& cd /tmp/librdkafka/ \
&& ./configure --prefix /usr \
&& make \
&& make install

WORKDIR /app

# 提取执行文件
COPY --from=0 /app/ipaas-job-reg ./

EXPOSE 9213
CMD ["./ipaas-job-reg"]

#源镜像
#FROM golang:1.12.4-alpine as kafka
#
## 设置工程配置文件的环境变量
#ENV PKG_CONFIG_PATH /usr/lib/pkgconfig
#
## 安装系统级依赖
#RUN apk add --no-cache git gcc musl-dev mercurial bash gcc g++ make pkgconfig openssl-dev
#
## 下载kafka
#RUN git clone https://github.com/edenhill/librdkafka.git $GOPATH/librdkafka
#
#WORKDIR $GOPATH/librdkafka
#RUN ./configure --prefix /usr  && \
#make && \
#make install
#
##################################################################################
##源镜像
#FROM golang:1.12.4-alpine as pods
#
#COPY --from=kafka /usr/lib/pkgconfig /usr/lib/pkgconfig
#COPY --from=kafka /go/librdkafka /go/librdkafka
#
## 安装系统级依赖
#RUN apk add --no-cache git bash openssl-dev
#
## 设置工程配置文件的环境变量
#ENV PKG_CONFIG_PATH /usr/lib/pkgconfig
#ENV MQTT_HOME $GOPATH/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources
#ENV BM_KAFKA_CONF_HOME $GOPATH/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources/resource/kafkaconfig.json
#ENV GO111MODULE on
#ENV LOGGER_USER "Alex"
#ENV LOGGER_DEBUG "true"
#ENV LOG_PATH /mqtt.log
#ENV GOPROXY https://goproxy.io
#
## 下载依赖
#
#RUN git clone https://github.com/PharbersDeveloper/MQTTMessageStorage.git  $GOPATH/src/github.com/PharbersDeveloper/MQTTMessageStorage
#
## 构建可执行文件
#RUN cd $GOPATH/src/github.com/PharbersDeveloper/MQTTMessageStorage && \
#    go build && go install
#
#
##################################################################################
##源镜像
#FROM golang:1.12.4-alpine
#
##作者
#MAINTAINER Pharbers "pqian@pharbers.com"
#
## 以LABEL行的变动(version的变动)来划分(变动以上)使用cache和(变动以下)不使用cache
#LABEL MQTTPods.version="0.0.3" maintainer="Alex"
#
## 设置工程配置文件的环境变量
#ENV PKG_CONFIG_PATH /usr/lib/pkgconfig
#ENV MQTT_HOME $GOPATH/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources
#ENV BM_KAFKA_CONF_HOME $GOPATH/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources/resource/kafkaconfig.json
#ENV GO111MODULE on
#ENV LOGGER_USER "Alex"
#ENV LOGGER_DEBUG "true"
#ENV LOG_PATH /mqtt.log
#
#COPY --from=kafka /usr/lib/pkgconfig /usr/lib/pkgconfig
#COPY --from=pods /go/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources/resource/kafkaconfig.json /go/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources/resource/kafkaconfig.json
#COPY --from=pods /go/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources /go/src/github.com/PharbersDeveloper/MQTTMessageStorage/resources
#COPY --from=pods ./go/bin/MQTTStorage ./go/bin/MQTTStorage
#
## 暴露端口
#EXPOSE 6542
#WORKDIR $GOPATH/bin
#
#ENTRYPOINT ["MQTTStorage"]
