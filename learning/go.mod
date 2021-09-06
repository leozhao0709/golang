module github.com/leozhao0709/learning

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.8.3
	github.com/golang/protobuf v1.4.2
	github.com/gomodule/redigo v1.8.1
	github.com/kr/text v0.2.0 // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.3.0
	github.com/labstack/gommon v0.3.0
	github.com/leozhao0709/musings v0.0.0-20201024010016-8fac21d7d6ca
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.15.0
	github.com/onsi/gomega v1.10.5
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.7.1
	github.com/streadway/amqp v1.0.0
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)

replace github.com/leozhao0709/musings => ../../musings
