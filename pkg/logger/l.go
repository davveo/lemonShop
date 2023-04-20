package logger

import (
	"fmt"
	"github.com/davveo/lemonShop/conf"
	"github.com/davveo/lemonShop/pkg/timeutil"
	"go.uber.org/zap"
	"os"
)

var GLogger *zap.Logger

func Init() (*zap.Logger, error) {
	lg, err := NewJSONLogger(
		WithDisableConsole(),
		WithField("domain", fmt.Sprintf("%s[%s]",
			conf.Conf.AppName, os.Getenv("active"))),
		WithTimeLayout(timeutil.CSTLayout),
		WithFileP(conf.Conf.Log.LogSavePath),
	)
	if err != nil {
		return nil, err
	}
	GLogger = lg
	return lg, nil
}
