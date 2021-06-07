package cmd

import (
	"strings"

	formatter "github.com/lacasian/logrus-module-formatter"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initLogging() {
	logging := viper.GetString("logging")

	if verbose {
		logging = "*=debug"
	}

	if vverbose {
		logging = "*=trace"
	}

	if logging == "" {
		logging = "*=info"
	}

	viper.Set("logging", logging)

	f, err := formatter.New(formatter.NewModulesMap(logging))
	if err != nil {
		panic(err)
	}

	logrus.SetFormatter(f)

	log.Debug("Debug mode")
}

func mustGetSubconfig(v *viper.Viper, key string, out interface{}) {
	err := unmarshalSubconfig(v, key, out)
	if err != nil {
		log.Fatal(err)
	}
}

func unmarshalSubconfig(v *viper.Viper, key string, out interface{}) error {
	vc := subtree(v, key)
	if vc == nil {
		return errors.Errorf("key '%s' not found", key)
	}
	err := vc.Unmarshal(out)
	return err
}

func subtree(v *viper.Viper, name string) *viper.Viper {
	r := viper.New()
	for _, key := range v.AllKeys() {
		if strings.Index(key, name+".") == 0 {
			r.Set(key[len(name)+1:], v.Get(key))
		}
	}
	return r
}
