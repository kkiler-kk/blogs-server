package routes

import (
	"kkblogs/blogs-api/controller/app"
	"kkblogs/blogs-api/routes/middleware"

	// "kkblogs/blogs-api/routes/middleware"

	"github.com/gin-gonic/gin"
)

func AppRoute(router *gin.RouterGroup) {
	appAuth := middleware.AppAuth
	helloController := app.Hello{}
	{
		router.GET("hello", helloController.Hello)
	}
	loginController := app.Login{}
	{
		router.POST("login", loginController.Login)  // 登录
		router.GET("logout", loginController.Logout) // 退出登录
	}
	messageRouter := app.Message{}
	{
		router.GET("ws", messageRouter.WsHandle)
	}
	uploadRouter := app.Upload{}
	{
		router.POST("upload", uploadRouter.ImagesUpload) // 图片文件上传
	}
	registerController := app.Register{}
	{
		router.POST("register", registerController.Register) // 注册
	}
	// 标签路由
	tagRouter := router.Group("tags")
	{
		tagController := app.Tag{}
		{
			tagRouter.GET("hot", tagController.Hot)             // 最热标签
			tagRouter.GET("", tagController.ListTags)           // 最热标签
			tagRouter.GET("detail", tagController.Detail)       // 查询所有标签
			tagRouter.GET("detail/:id", tagController.DetailId) // 查询标签 根据id
		}
	}
	// 收藏
	goodRouter := router.Group("collect")
	{
		goodController := app.Good{}
		{
			goodRouter.POST("/user/:id", goodController.GetGoodByUserId) // 根据用户id查询他的收藏
			goodRouter.POST("/is/collect", goodController.IsGood)        // 根据文章id 用户id 查看是否有没有收藏此文章
		}
	}
	// 点赞
	likeRouter := router.Group("like")
	{
		likeController := app.Like{}
		{
			likeRouter.POST("/is/like", likeController.IsLike) // 根据文章id 用户id 查看是否有没有收藏此喜欢
		}
	}
	// 用户
	usersRouter := router.Group("users")
	{
		usersController := app.Users{}
		{
			usersRouter.GET("currentUser", usersController.CurrentUser)        // 获取用户信息
			usersRouter.POST(":id", usersController.GetUserById)               // 根据id获取用户信息
			usersRouter.POST("info/:id", usersController.GetUserInfo)          // 获取用户详细信息
			usersRouter.POST("change/pwd", appAuth, usersController.ChangePwd) // 修改密码
			usersRouter.POST("change/info", appAuth, usersController.EditInfo) // 修改信息
			usersRouter.POST("article/:id", usersController.MyArticle)         // 查询用户文章
		}
	}
	// 关注
	followRouter := router.Group("follow")
	{
		followController := app.Follow{}
		{
			followRouter.POST("add", appAuth, followController.AddFollow)       // 添加关注
			followRouter.POST("remove", appAuth, followController.RemoveFollow) // 取消关注
			followRouter.POST(":id", followController.GetUserFollow)            // 获取关注数 粉丝数
			followRouter.POST("is/follow", followController.IsFollow)           // 查看是否被关注
			followRouter.POST("follow/:id", followController.MyFollow)          // 查看关注
			followRouter.POST("fans/:id", followController.MyFans)              // 查看粉丝
		}
	}
	// 文章路由
	articlesRouter := router.Group("articles")
	{
		articlesController := app.Articles{}
		{
			articlesRouter.POST("", articlesController.ListArticles)             // 查询文章列表 时间排序
			articlesRouter.POST("hot", articlesController.HotArticles)           // 最热文章列表（title）
			articlesRouter.POST("new", articlesController.NewArticles)           // 最新文章列表（title）
			articlesRouter.POST("listArchives", articlesController.ListArchives) // 文章归档
			articlesRouter.POST("view/:id", articlesController.ArticlesView)     // 文章详情
			articlesRouter.POST("publish", appAuth, articlesController.Publish)  // 发布文章
			articlesRouter.POST(":id", articlesController.ArticlesView)          // 文章详情
			articlesRouter.POST("search", articlesController.Search)             // 搜索文章
		}
	}
	// 分类
	categorysRouter := router.Group("categorys")
	{
		categorysController := app.Category{}
		{
			categorysRouter.GET("", categorysController.ListCategory)       // 查询所有分类
			categorysRouter.GET("detail", categorysController.Detail)       // 查询所有文章分类
			categorysRouter.GET("detail/:id", categorysController.DetailId) // 查询所有文章分类
		}
	}
	// 评论
	commentRouter := router.Group("comments")
	{
		commentController := app.Comments{}
		{
			commentRouter.GET("article/:id", commentController.ArticlesCommentId)         // 根据文章id 查看文章详情评论
			commentRouter.POST("create/change", appAuth, commentController.CreateComment) // 创建评论
		}
	}

}
