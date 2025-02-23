package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct{
	AmqpUrl string `yaml:"amqp_url"`
	MongoUrl string `yaml:"mongo_url"`
	RpcQueueName string `yaml:"rpc_que_name"`
	Port uint32 `yaml:"port"`
	Host string `yaml:"host"`
}

func Load(path string) (*Config, error) {
	def := &Config{
		Port: 50054,
		MongoUrl: "mongodb://localhost:27017",
		AmqpUrl: "amqp://guest:guest@localhost:5672/",
		RpcQueueName: "basket",
		Host: "localhost",
	}

	file, err := os.Open(path)
	if err != nil{
		return def, nil
	}

	body, err := io.ReadAll(file)
	if err != nil{
		return def, nil
	}

	var cfg Config

	err = yaml.Unmarshal(body, &cfg)
	return &cfg, err
}