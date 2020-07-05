package conf

import (
	"os"
	"strings"
	"time"

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
		TimeoutDuration time.Duration
		Port            string
		Debug           bool
		DevImagesPath   string
		ImageHeight     int
		ImageWidth      int
	}
	Gcs struct {
		BucketName string
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

	C.Sv.DevImagesPath = dir + "/" + C.Sv.DevImagesPath
	C.Sv.TimeoutDuration = time.Duration(C.Sv.Timeout) * time.Second

	if C.Sv.Debug {
		if _, err := os.Stat(C.Sv.DevImagesPath); os.IsNotExist(err) {
			if err := os.Mkdir(C.Sv.DevImagesPath, 0755); err != nil {
				panic(err)
			}
		}

	}

	spew.Dump(C)
}
