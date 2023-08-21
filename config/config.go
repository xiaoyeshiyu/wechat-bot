package config

import (
	"os"

	"github.com/spf13/viper"
)

// Configuration 项目配置，内容较少，平铺即可
type Configuration struct {
	ApiKey   string `mapstructure:"api_key"`   // chatGPT key
	AutoPass bool   `mapstructure:"auto_pass"` // 添加好友是否自动通过
	ChatUrl  string `mapstructure:"chat_url"`  // chatGPT url，可使用自己的url
	Log      Log    `mapstructure:"log"`
}

// Log 日志配置
type Log struct {
	File    string `mapstructure:"file,omitempty"`
	Level   string `mapstructure:"level"`
	MaxSize int    `mapstructure:"maxSize"`
	Backups int    `mapstructure:"backups"`
	MaxAge  int    `mapstructure:"maxAge"`
}

var config Configuration

// LoadConfig 加载配置
func LoadConfig() error {
	configFile := "./config.json"

	// 绑定配置文件
	// todo 增加读取命令行参数和环境变量的形式
	configViper := viper.New()
	configViper.SetConfigFile(configFile)
	err := configViper.ReadInConfig()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	return configViper.Unmarshal(&config)
}

func GetApiKey() string {
	// 个人Key，用作默认key
	if config.ApiKey == "" {
		return "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"
	}
	return config.ApiKey
}

func GetAutoPass() bool {
	return config.AutoPass
}

func GetChatUrl() string {
	// 没有指定url，则使用openai默认url
	if config.ChatUrl == "" {
		return "https://api.openai.com/"
	}
	return config.ChatUrl
}

func GetLogFile() string {
	if config.Log.File == "" {
		return "./wechat_bot.log"
	}
	return config.Log.File
}
func GetLogLevel() string {
	if config.Log.Level == "" {
		return "info"
	}
	return config.Log.Level
}

func GetLogMaxSize() int {
	if config.Log.MaxSize == 0 {
		return 100
	}
	return config.Log.MaxSize
}

func GetLogBackups() int {
	if config.Log.Backups == 0 {
		return 3
	}
	return config.Log.Backups
}

func GetLogMaxAge() int {
	if config.Log.MaxAge == 0 {
		return 30
	}
	return config.Log.MaxAge
}
