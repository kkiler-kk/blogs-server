package config

import "gorm.io/gorm/logger"

type MySQL struct {
	Username      string          `ini:"username"`
	Password      string          `ini:"password"`
	Host          string          `ini:"host"`
	Port          int             `ini:"port"`
	DbName        string          `ini:"dbName"`
	Prefix        string          `ini:"prefix"`
	Extend        string          `ini:"extend"`
	SingularTable bool            `ini:"singularTable"`
	LogLevel      logger.LogLevel `ini:"logLevel"`
	MaxIdleConns  int             `ini:"maxIdleConns"`
	MaxOpenConns  int             `ini:"maxOpenConns"`
	LogColorful   bool            `ini:"logColorful"`
}
