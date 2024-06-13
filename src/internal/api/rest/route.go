package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.ruoyi.com/src/config"
	"go.ruoyi.com/src/internal/middleware"
	"log"
)

func InitHttp() {
	r := gin.New()
	r.HandleMethodNotAllowed = true
	r.MaxMultipartMemory = 20 << 20

	r.Use(middleware.Cors())
	//r.Use(middleware.TokenMiddleware()) //token检测
	//r.Use(middleware.DataEncrypr())                //解密中间件，将请求体解密给日志存放了reqbody参数
	r.Use(middleware.LoggerMiddleWare()) //日志捕捉
	//r.Use(middleware.GlobalErrorMiddleware())      //错误捕捉
	r.Use(middleware.UnifiedResponseMiddleware())  //全局统一返回格式，添加了rsa
	r.NoRoute(middleware.NotFoundHandler)          //404
	r.NoMethod(middleware.MethodNotAllowedHandler) //405，方法为找到
	initApi(r)                                     //挂载请求路径
	fmt.Println("中间加载结束", config.Port)
	//err := r.RunTLS(config.Port, "../../config/https/certificate.crt", "../../config/https/private.key")
	err := r.Run(config.Port)
	//log.Println(su)
	if err != nil {
		log.Print(err)
		panic("端口启动错误")
	}
}
func initApi(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.PATCH("/patch", func(c *gin.Context) {
			c.Set("res", "success")
		})
		api.GET("/test", func(c *gin.Context) {
			log.Println("请求到")

			//c.JSON(http.StatusOK, gin.H{"message": "success"})
			c.Set("res", "success")
			//return
		})
		//auth := api.Group("/auth")
		//{
		//	auth.POST("/login", auth2.Login)
		//}
		//ws := api.Group("/v1/ws")
		//test := api.Group("/v1/test")
		//{
		//	test.GET("/ping", func(context *gin.Context) {
		//		context.Set("res", "success")
		//		return
		//	})
		//
		//}
	}
	//ws:=r.Group("/v1/api/ws")
	//static:=r.Group("/v1/api/static")

}
