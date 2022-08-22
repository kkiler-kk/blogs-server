package config

type Redis struct {
	Addr     string `ini:"addr"`
	Password string `ini:"password"`
	DB       int    `ini:"db"`
}
