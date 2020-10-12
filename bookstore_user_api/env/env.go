package env

import (
	"flag"

	"github.com/labstack/gommon/log"
)

var currentEnv string

func init() {
	var environment = flag.String("env", "dev", "which env here?(prod/dev/test)")
	flag.Parse()
	flag.PrintDefaults()
	setCurrentEnv(environment)
}

func setCurrentEnv(env *string) {
	currentEnv = *env
	log.Printf("Set current env as %s", currentEnv)
}

// GetCurrentEnv get current env
func GetCurrentEnv() string {
	return currentEnv
}
