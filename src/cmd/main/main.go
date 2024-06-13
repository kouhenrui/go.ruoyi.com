package main

import (
	"go.ruoyi.com/src/config"
	"go.ruoyi.com/src/internal/api/rest"
)

func main() {
	config.InitConfig()
	rest.InitHttp()

}
