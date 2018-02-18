package env

import (
	config "algtmapi/appobject/configobject"
	"algtmapi/assets"
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/deadcheat/goacors"
	"github.com/spf13/viper"
)

// 環境設定として保持する情報の構造体
var (
	// デフォルト環境
	OnDevelopment  = false
	configfilename = "local"

	// アプリケーション設定
	Server   *config.ServerConfig
	CorsConf goacors.GoaCORSConfig

	// Time設定
	DataTimeLayout string
)

// init パッケージの初期化処理
func init() {
	// 初期設定
	initialize()
	InitializeEnv()
}

// デフォルト値設定
const (
	defaultTimeLocation   = "Asia/Tokyo"
	defaultTimeZoneOffset = 9 * 60 * 60
	defaultDataTimeLayout = "2006-01-02 15:04:05"
)

// InitializeEnv パッケージの初期化処理
func InitializeEnv() {
	setTomlConfig()
	setTimeConfig()
}

func setTomlConfig() {
	// 設定ファイル
	confDir := "config"
	fileName := configfilename
	ext := "toml"
	filePath := fmt.Sprintf("%s/%s.%s", confDir, fileName, ext)

	// Asset経由でファイルを取得しに行く
	viper.SetConfigType(ext)
	data, err := assets.Asset(filePath)
	if err != nil {
		log.Panic(err, filePath)
	}

	// go-bindataのAsset、およびconfigファイルのbindataが存在する場合
	viper.ReadConfig(bytes.NewBuffer(data))
	// 設定読み込み
	_ = viper.UnmarshalKey("server", &Server)
	_ = viper.UnmarshalKey("cors", &CorsConf)
}

func setTimeConfig() {
	loc, err := time.LoadLocation(defaultTimeLocation)
	if err != nil {
		loc = time.FixedZone(defaultTimeLocation, defaultTimeZoneOffset)
	}
	time.Local = loc
	DataTimeLayout = defaultDataTimeLayout
}
