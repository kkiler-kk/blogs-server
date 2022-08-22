package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"kkblogs/blogs-api/common"
	"kkblogs/blogs-api/config"
	"kkblogs/blogs-api/routes"
	"kkblogs/blogs-api/routes/middleware/handler"
	"log"
)

func main() {
	// 初始化项目
	closes := initApp()
	//关闭资源
	for _, fn := range closes {
		defer fn()
	}
	//路由
	engine := gin.Default()
	// 全局统一异常处理
	engine.Use(handler.Recover)
	// 设置路由
	routes.SetupRouter(engine)

	// 新建一个定时任务对象
	// 根据cron表达式进行时间调度，cron可以精确到秒，大部分表达式格式也是从秒开始。
	//crontab := cron.New()  默认从分开始进行时间调度
	crontab := cron.New(cron.WithSeconds()) //精确到秒
	//定义定时器调用的任务函数
	task1 := common.SyncDataBaseComment
	task2 := common.SyncDataBaseView
	//定时任务
	spec := "0 0 1 * * ?" //cron表达式，每五秒一次
	// 添加定时任务,
	crontab.AddFunc(spec, task1)
	crontab.AddFunc(spec, task2)
	// 启动定时器
	crontab.Start()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
	//select {} //阻塞主线程停止
	// 运行
	if err := engine.Run(config.Config.Server.Addr); err != nil {
		log.Println("运行错误")
		return
	}

}
