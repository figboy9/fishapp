package conf

import (
	"os"
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
	Sv struct {
		Timeout         int64
		Port            string
		Debug           bool
		DefaultPageSize int64
		ImageChunkSize  int64
	}
	Nats struct {
		URL        string
		ClusterID  string
		QueueGroup string
	}
	API struct {
		ImageURL string `mapstructure:"image_url"`
	}
}

var C config

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath(dir + "/conf")

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
