package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

const (
	kAppName       = "APP_NAME"
	kConfigServer  = "CONFIG_SERVER"
	kConfigLabel   = "CONFIG_LABEL"
	kConfigProfile = "CONFIG_PROFILE"
	kConfigType    = "CONFIG_TYPE"
)

var (
	App AppConfig
)

type AppConfig struct {
	Name            string
	SearchUrl       string `mapstructure:"search_url"`
	JwtSecret       string
	PageSize        int
	PrefixUrl       string
	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	Server   Server
	Redis    Redis
	Database Database
	AliOss   AliOss
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type AliOss struct {
	AliyunEndPoint        string
	AliyunAccessKeyId     string
	AliyunAccessKeySecret string
	AliyunBucketName      string
	AliyunPrefix          string
	AliyunDomain          string
}

func Setup() {
	viper.AutomaticEnv()
	initDefault()
}

func init() {
	viper.AutomaticEnv()
	initDefault()

	if err := loadRemoteConfig(); err != nil {
		log.Fatal("Fail to load config", err)
	}

	if err := sub("app", &App); err != nil {
		log.Fatal("Fail to parse config", err)
	}
}

func initDefault() {
	viper.SetDefault(kAppName, "go-app")
	viper.SetDefault(kConfigServer, "http://localhost:8888")
	viper.SetDefault(kConfigLabel, "master")
	viper.SetDefault(kConfigProfile, "dev")
	viper.SetDefault(kConfigType, "yml")
}

func loadRemoteConfig() (err error) {
	confAddr := fmt.Sprintf("%v/%v/%v-%v.%v",
		viper.Get(kConfigServer), viper.Get(kConfigLabel),
		viper.Get(kAppName), viper.Get(kConfigProfile),
		viper.Get(kConfigType))
	resp, err := http.Get(confAddr)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	viper.SetConfigType(viper.GetString(kConfigType))
	if err = viper.ReadConfig(resp.Body); err != nil {
		return
	}
	log.Println("Load config from: ", confAddr)
	return
}

func sub(key string, value interface{}) error {
	sub := viper.Sub(key)
	sub.AutomaticEnv()
	sub.SetEnvPrefix(key)
	return sub.Unmarshal(value)
}
