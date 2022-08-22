package config

type Rabbitmq struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
}
