package main

import (
	"git.oa00.com/go/logger"
	"git.oa00.com/go/mysql"
	"git.oa00.com/go/redis"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"kkblogs/blogs-api/config"
	"kkblogs/blogs-api/lib/bean"
	"log"
	"reflect"
)

func initApp() (closes []func()) {
	// 关闭资源f
	closes = []func(){}
	// 初始化配置文件
	config.InitConfig(&config.IniConfig{
		ConfigPath:       "config",
		RunModelErrAllow: true,
	})

	// 系统分页
	bean.InitPage(config.Config.Page.MinLimit, config.Config.Page.MaxLimit, config.Config.Page.DefaultLimit)
	logLevel := logger.InfoLevel
	if config.Config.Debug {
		logLevel = logger.DebugLevel
	}
	// 初始化日志
	logger.InitLogger(&logger.LoggerConfig{
		Director:      "log",
		Level:         logLevel,
		ShowLine:      true,
		StacktraceKey: "",
		LinkName:      "",
		LogInConsole:  true,
		EncodeLevel:   logger.LowercaseColorLevelEncoder,
		Prefix:        "",
	})

	// 初始化redis
	pong, err := redis.InitRedis(&redis.RedisConfig{
		Addr:     config.Config.Redis.Addr,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DB,
	})
	if err != nil {
		log.Fatalln("redis connect faild", err)
	} else {
		log.Println("redis pong ", pong)
	}
	// 初始化七牛云

	// 初始化数据库
	if err := mysql.InitMysql(&mysql.DbConfig{
		Username:      config.Config.Mysql.Username,
		Password:      config.Config.Mysql.Password,
		Host:          config.Config.Mysql.Host,
		Port:          config.Config.Mysql.Port,
		DbName:        config.Config.Mysql.DbName,
		Prefix:        config.Config.Mysql.Prefix,
		Extend:        config.Config.Mysql.Extend,
		SingularTable: config.Config.Mysql.SingularTable,
		LogColorful:   config.Config.Mysql.LogColorful,
		LogLevel:      config.Config.Mysql.LogLevel,
		MaxIdleConns:  config.Config.Mysql.MaxIdleConns,
		MaxOpenConns:  config.Config.Mysql.MaxOpenConns,
	}); err != nil {
		log.Panicln(err)
		return
	}
	// 初始化线程池
	//thread.NewWorkerPool(20)
	// 验证器处理
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})
	}
	return
}
