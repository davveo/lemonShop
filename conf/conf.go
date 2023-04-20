package conf

import (
	"github.com/davveo/lemonShop/pkg/config"
	"os"
)

var Conf AppConf

const (
	activeProfiles = "active"
)

func Init() (*AppConf, error) {
	err := config.Unmarshal(ConfYamlDir,
		&Conf, config.ProfileFunc(func() string {
			return os.Getenv(activeProfiles)
		}))
	if err != nil {
		return nil, err
	}
	return &Conf, nil
}
