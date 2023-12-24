package main

import (
	"flag"
	"os"

	"github.com/aws-samples/aws-global-accelerator-custom-routing-workshop/echo-cli/client"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

type ClientConfig struct {
	Open bool `yaml:"open"`
	Port int  `yaml:"port"`
}

type Config struct {
	// config for log
	Log struct {
		Level  string `yaml:"level"`
		Format string `yaml:"format"`
	} `yaml:"log"`
	// config for server
	Server struct {
		Host      string `yaml:"host"`
		TestCount int    `yaml:"test_count"`
	} `yaml:"server"`
	// config for client
	Client struct {
		// config for tcp
		Tcp ClientConfig `yaml:"tcp"`
		// config for udp
		Udp ClientConfig `yaml:"udp"`
		// config for http
		Http ClientConfig `yaml:"http"`
		// config for grpc
		Grpc ClientConfig `yaml:"grpc"`
		// config for websocket
		Websocket ClientConfig `yaml:"websocket"`
	} `yaml:"client"`
}

func main() {
	configFileName := flag.String("f", "config.yaml", "config file name")
	flag.Parse()
	config := Config{}
	// get configyaml from file config.yaml
	configYaml, err := os.ReadFile(*configFileName)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configYaml, &config)
	if err != nil {
		panic(err)
	}
	// set log format and level
	logLevel := config.Log.Level
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	if config.Log.Format == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
	logrus.SetLevel(level)
	logrus.SetOutput(os.Stdout)
	logrus.Debugf("%+v", config)

	// start tcp test
	if config.Client.Tcp.Open {
		client.StartTcpTest(config.Server.Host, config.Client.Tcp.Port, config.Server.TestCount)
	}

	// start udp test
	if config.Client.Udp.Open {
		client.StartUdpTest(config.Server.Host, config.Client.Udp.Port, config.Server.TestCount)
	}

	// start http test
	if config.Client.Http.Open {
		client.StartHttpTest(config.Server.Host, config.Client.Http.Port, config.Server.TestCount)
	}

	// start websocket test
	if config.Client.Websocket.Open {
		client.StartWebsocketTest(config.Server.Host, config.Client.Websocket.Port, config.Server.TestCount)
	}

	// start grpc test
	if config.Client.Grpc.Open {
		client.StartGrpcTest(config.Server.Host, config.Client.Grpc.Port, config.Server.TestCount)
	}
}
