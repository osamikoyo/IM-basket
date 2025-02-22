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
	file, err := os.Open(path)
	if err != nil{
		return nil, err
	}

	body, err := io.ReadAll(file)
	if err != nil{
		return nil, err
	}

	var cfg Config

	err = yaml.Unmarshal(body, &cfg)
	return &cfg, err
}