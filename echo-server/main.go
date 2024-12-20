package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/aws-samples/aws-global-accelerator-custom-routing-workshop/echo-server/service"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Open  bool   `yaml:"open"`
	Host  string `yaml:"host"`
	Ports []int  `yaml:"ports"`
}

type Config struct {
	// config for log
	Log struct {
		Level  string `yaml:"level"`
		Format string `yaml:"format"`
	} `yaml:"log"`
	// config for server
	Server struct {
		// config for tcp
		Tcp ServerConfig `yaml:"tcp"`
		// config for udp
		Udp ServerConfig `yaml:"udp"`
		// config for http
		Http ServerConfig `yaml:"http"`
		// config for websocket
		Websocket ServerConfig `yaml:"websocket"`
		// config for grpc
		Grpc ServerConfig `yaml:"grpc"`
	} `yaml:"server"`
}

func loadConfig(configFileName string) Config {
	config := Config{}
	// get configyaml from file config.yaml
	configYaml, err := os.ReadFile(configFileName)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configYaml, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func initLog(config Config) {
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
}

func main() {
	configFileName := flag.String("f", "config.yaml", "config file name")
	flag.Parse()
	config := loadConfig(*configFileName)
	initLog(config)

	// start tcp server
	if config.Server.Tcp.Open {
		for _, port := range config.Server.Tcp.Ports {
			tcpserver, _ := service.NewTcpServer(config.Server.Tcp.Host, port)
			go tcpserver.Start()
			logrus.Info("tcp server started, listening on port ", port)
		}
	}

	// start udp server
	if config.Server.Udp.Open {
		for _, port := range config.Server.Udp.Ports {
			udpserver, _ := service.NewUdpServer(config.Server.Udp.Host, port)
			go udpserver.Start()
			logrus.Info("udp server started, listening on port ", port)
		}
	}

	// start http server
	if config.Server.Http.Open {
		service.RegisterHttpHandler()
		for _, port := range config.Server.Http.Ports {
			httpserver, _ := service.NewHttpServer(config.Server.Http.Host, port)
			go httpserver.Start()
			logrus.Info("http server started, listening on port ", port)
		}
	}

	// start websocket server
	if config.Server.Websocket.Open {
		for _, port := range config.Server.Websocket.Ports {
			websocketserver, _ := service.NewWebSocketServer(config.Server.Websocket.Host, port)
			go websocketserver.Start()
			logrus.Info("websocket server started, listening on port ", port)
		}
	}

	// start grpc server
	if config.Server.Grpc.Open {
		for _, port := range config.Server.Grpc.Ports {
			grpcserver, _ := service.NewGrpcServer(config.Server.Grpc.Host, port)
			go grpcserver.Start()
			logrus.Info("grpc server started, listening on port ", port)
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	<-c
}
