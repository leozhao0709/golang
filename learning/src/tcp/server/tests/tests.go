package tests

import (
	"runtime/debug"

	"github.com/labstack/gommon/log"
)

func Test() {
	log.Debugf("debug.....%v\n", string(debug.Stack()))
}
