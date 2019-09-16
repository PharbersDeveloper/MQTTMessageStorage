package env

import (
	"os"
)

const (
	//Project env key
	ProjectName = "PROJECT_NAME"

	//Log env key
	LogTimeFormat = "BP_LOG_TIME_FORMAT"
	LogOutput     = "BP_LOG_OUTPUT"
	LogLevel      = "BP_LOG_LEVEL"

	//kafka env key
	KafkaConfigEnable = "BP_KAFKA_CONFIG_ENABLE"
	KafkaConfigPath = "BP_KAFKA_CONFIG_PATH"
)

func SetEnv() {
	//项目范围内的环境变量
	_ = os.Setenv(ProjectName, "MQTTMessageStorage")

	//log
	_ = os.Setenv(LogTimeFormat, "2006-01-02 15:04:05")
	_ = os.Setenv(LogOutput, "console")
	//_ = os.Setenv(env.LogOutput, "/Users/qianpeng/bplogs/mqtt-message-storage.log")
	_ = os.Setenv(LogLevel, "info")

	//kafka
	//_ = os.Setenv(KafkaConfigPath, "../resources/kafka_config.json")
}
