/******************************************************
# DESC    : env var & configure
# AUTHOR  : Alex Stocks
# VERSION : 1.0
# LICENCE : Apache License 2.0
# EMAIL   : alexstocks@foxmail.com
# MOD     : 2016-07-01 15:20
# FILE    : config.go
******************************************************/

package main

import (
	"fmt"
	"os"
	"path"
)

import (
	"github.com/AlexStocks/goext/log"
	log "github.com/AlexStocks/log4go"
	config "github.com/koding/multiconfig"
)

import (
	"github.com/AlexStocks/dubbogo/common"
	"github.com/AlexStocks/dubbogo/registry"
)

const (
	APP_CONF_FILE     = "APP_CONF_FILE"
	APP_LOG_CONF_FILE = "APP_LOG_CONF_FILE"
)

var (
	conf *ClientConfig
)

type (
	// Client holds supported types by the multiconfig package
	ClientConfig struct {
		// pprof
		Pprof_Enabled bool `default:"false"`
		Pprof_Port    int  `default:"10086"`

		// client
		Request_Timeout string `default:"5s"` // 500ms, 1m
		NET_IO_Timeout  string `default:"5s"` // 500ms, 1m
		Retries         int    `default:"1"`
		// Connection Pool
		Pool_Size       int    `default:"128"`
		Pool_TTL        string `default:"1m"`
		Connect_Timeout string `default:"100ms"`

		// codec & selector & transport & registry
		Selector     string `default:"cache"`
		Selector_TTL string `default:"10m"`
		Registry     string `default:"zookeeper"`
		// application
		Application_Config common.ApplicationConfig
		Registry_Config    registry.RegistryConfig
		// 一个客户端只允许使用一个service的其中一个group和其中一个version
		Service_List []registry.ServiceConfig
	}
)

func initClientConfig() error {
	var (
		confFile string
	)

	// configure
	confFile = os.Getenv(APP_CONF_FILE)
	if confFile == "" {
		panic(fmt.Sprintf("application configure file name is nil"))
		return nil // I know it is of no usage. Just Err Protection.
	}
	if path.Ext(confFile) != ".toml" {
		panic(fmt.Sprintf("application configure file name{%v} suffix must be .toml", confFile))
		return nil
	}
	conf = new(ClientConfig)
	config.MustLoadWithPath(confFile, conf)
	gxlog.CInfo("config{%#v}\n", conf)

	// log
	confFile = os.Getenv(APP_LOG_CONF_FILE)
	if confFile == "" {
		panic(fmt.Sprintf("log configure file name is nil"))
		return nil
	}
	if path.Ext(confFile) != ".xml" {
		panic(fmt.Sprintf("log configure file name{%v} suffix must be .xml", confFile))
		return nil
	}
	log.LoadConfiguration(confFile)

	return nil
}
