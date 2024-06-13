package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"go.ruoyi.com/src/internal/pojo"
	"log"
	"time"
)

var (
	v   *viper.Viper
	err error
)

var (
	Port            string
	Tcp             string
	InterceptPrefix string
	language        string
	JWTKey          string
	WhiteUrl        = []string{"/v1/api/auth/login", "/v1/api/auth/register", "/v1/api/public/.*", "/v1/api/monitor/.*", "/api/v1/test/ping", "/api/patch", "/api/test"}
	Captcha         pojo.Captcha
	mysqlClient     pojo.MysqlConf
	redisClient     pojo.RedisConf
	casbinClient    pojo.CabinConf
	logConf         pojo.LogCof //连接日志实例化参数
	mq              pojo.RabbitmqConf
	etcdArry        = []string{}
)

const (
	AdminExp = time.Duration(time.Hour * 24 * 5)
	UserExp  = time.Duration(time.Hour * 24 * 5)
)

func InitConfig() {
	log.Println("实例化配置文件")
	// 构建 Viper 实例
	v = viper.New()
	v.SetConfigFile("conf.yaml") // 指定配置文件路径
	//v.SetConfigFile("/data/goland/config")
	v.SetConfigName("conf")                                                // 配置文件名称(无扩展名)
	v.SetConfigType("yaml")                                                // 如果配置文件的名称中没有扩展名，则需要配置此项
	v.AddConfigPath("D:\\\\code_space\\\\go\\go.encrypt.gin\\src\\config") // 多次调用以添加多个搜索路径
	v.AddConfigPath("../../config/")                                       // 还可以在工作目录中查找配置
	viper.AddConfigPath(".")                                               // 设置配置文件和可执行二进制文件在用一个目录
	//v.AddConfigPath("/data/goland/") // 查找配置文件所在的路径,服务器容器目录
	// 查找并读取配置文件
	if err = v.ReadInConfig(); err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viperLoadConf()
	v.WatchConfig() //开启监听
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file updated.")
		viperLoadConf() // 加载配置的方法
	})

}

func viperLoadConf() {
	//读取单条配置文件
	Port = v.GetString("server.port")

	Tcp = v.GetString("tcp")
	language = v.GetString("language")
	JWTKey = v.GetString("jwt")
	InterceptPrefix = v.GetString("InterceptPrefix")
	//log.Println(Port, Tcp, language, "端口号")
	//日志路径及名称设置
	logConfig := v.GetStringMap("log")

	captcha := v.GetStringMap("captcha")
	mysql := v.GetStringMap("mysql") //读取MySQL配置
	redis := v.GetStringMap("redis") //读取redis配置
	//mq := v.GetStringMap("rabbitmq") //读取rabbitmq配置
	cn := v.GetStringMap("cabin") //读取casbin配置
	//ck := v.GetStringMap("click")    //读取click house配置

	rabbitmqClient := v.GetStringMap("rabbitmq")

	//map转struct
	mapstructure.Decode(mysql, &mysqlClient)
	mapstructure.Decode(redis, &redisClient)
	mapstructure.Decode(rabbitmqClient, &mq)
	mapstructure.Decode(logConfig, &logConf)
	mapstructure.Decode(cn, &casbinClient)
	//mapstructure.Decode(ck, &ClickConfig)

	mapstructure.Decode(captcha, &Captcha)
	Captcha.Expired = Captcha.Expired * time.Minute
	etcdConnect := v.GetStringSlice("etcd")
	//kafka := v.GetStringSlice("kafka")
	//oracle := v.GetStringSlice("oracle")
	etcdArry = append(etcdArry, etcdConnect...)
	//KafkaArry = append(KafkaArry, kafka...)
	log.Println(Tcp)
	log.Println("全局配置文件信息读取无误,开始载入")
	//Dbinit()         //mysql初始化
	//Redisinit() //redis初始化

	//translate.TranslateInit(language) //i18语言设置
	//logger.Loginit(&logConf) //日志初始化
	//etcd.EtcdInit(etcdArry)
	//db.Dbinit(&mysqlClient)          //mysql连接
	//rediss.Redisinit(&redisClient)   //redis连接
	//casbin.CasbinInit(&casbinClient) //casbin连接
	//etcd.EtcdInit(etcdConnect)
	//repository.InitAutoMigrate() //加载实体表
	//rabbitmq.Mqinit(&mq)
}
