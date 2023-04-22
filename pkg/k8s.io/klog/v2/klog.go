package klog

import (
	"fmt"

	"github.com/neilotoole/shelleditor/logging"
)

func Errorf(format string, args ...interface{}) {
	logging.Logger().Error(fmt.Sprintf(format, args...))
}

func ErrorDepth(d int, err error) {
	if err == nil {
		return
	}
	logging.Logger().Error(err.Error())
}
