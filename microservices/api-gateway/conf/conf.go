package conf

import (
	"log"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	Sv struct {
		Timeout        int64
		Port           string
		Debug          bool
		ChunkDataSize  int
		ClientTimezone string
		ClientLocation *time.Location
	}
	API struct {
		PostURL  string `mapstructure:"post_url"`
		UserURL  string `mapstructure:"user_url"`
		ChatURL  string `mapstructure:"chat_url"`
		ImageURL string `mapstructure:"image_url"`
	}
	Auth struct {
		PubJwtkey string `mapstructure:"pub_jwtkey"`
	}
	Graphql struct {
		URL        string
		Playground struct {
			Enable bool
			URL    string
			User   string
			Pass   string
		}
	}
	Nats struct {
		URL        string
		ClusterID  string
		QueueGroup string
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
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		log.Fatal(err)
	}

	l, err := time.LoadLocation(C.Sv.ClientTimezone)
	if err != nil {
		log.Fatal(err)
	}

	C.Sv.ClientLocation = l

	spew.Dump(C)

}
