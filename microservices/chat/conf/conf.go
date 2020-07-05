package conf

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	Db struct {
		Dbms                 string
		Name                 string
		User                 string
		Pass                 string
		Net                  string
		Host                 string
		Port                 string
		Parsetime            bool
		AllowNativePasswords bool
	}
	Kvs struct {
		Db       int
		Pass     string
		Host     string
		Port     string
		Net      string
		Sentinel struct {
			Host       string
			Port       string
			MasterName string
			Pass       string
		}
	}
	Nats struct {
		URL        string
		ClusterID  string
		QueueGroup string
	}
	Sv struct {
		Timeout        int64
		Port           string
		Debug          bool
		ImageChunkSize int64
	}
	API struct {
		ImageURL string `mapstructure:"image_url"`
	}
}

var C config

func init() {

	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath("conf")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		panic(err)
	}

	spew.Dump(C)
}
