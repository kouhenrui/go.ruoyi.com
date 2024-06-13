package pojo

import "time"

/**
 * @ClassName conf
 * @Description TODO
 * @Author khr
 * @Date 2023/7/29 14:39
 * @Version 1.0
 */
type MysqlConf struct {
	UserName string `json:"username,omitempty" yaml:"username" mapstructure:"username"`
	PassWord string `json:"password,omitempty" yaml:"password" mapstructure:"password"`
	HOST     string `json:"host,omitempty" yaml:"host" mapstructure:"host"`
	Port     string `json:"port,omitempty" yaml:"port" mapstructure:"port"`
	DATABASE string `json:"database,omitempty" yaml:"database" mapstructure:"database"`
	CHARSET  string `json:"charset,omitempty" yaml:"charset" mapstructure:"charset"`
	TimeOut  int64  `json:"timeout,omitempty" yaml:"timeout" mapstructure:"timeout"`
}
type RedisConf struct {
	UserName   string `json:"username,omitempty" yaml:"username"`
	PassWord   string `json:"password,omitempty" yaml:"password"`
	Host       string `json:"host,omitempty" yaml:"host"`
	Port       string `json:"port,omitempty" yaml:"port"`
	Db         int    `json:"db,omitempty" yaml:"db"`
	PoolSize   int    `json:"poolsize,omitempty" yaml:"poolsize"`
	MaxRetries int    `json:"maxRetries" yaml:"maxRetries"`
}
type RabbitmqConf struct {
	Url      string `json:"url,omitempty" yaml:"url" `
	UserName string `json:"username" yaml:"username"`
	PassWord string `json:"password" yaml:"password"`
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
}
type CabinConf struct {
	Type     string `json:"type" yaml:"type" mapstructure:"type"`
	UserName string `json:"username,omitempty" yaml:"username" mapstructure:"username"`
	PassWord string `json:"password,omitempty" yaml:"password" mapstructure:"password"`
	HOST     string `json:"host,omitempty" yaml:"host" mapstructure:"host"`
	Port     string `json:"port,omitempty" yaml:"port" mapstructure:"port"`
	DATABASE string `json:"database,omitempty" yaml:"database" mapstructure:"database"`
	Exist    bool   `json:"exist,omitempty" yaml:"exist" mapstructure:"exist"`
}

type EtcdConf struct {
	Host string `json:"host,omitempty" yaml:"host"`
	Port string `json:"port,omitempty" yaml:"port"`
}
type ClickConf struct {
	Host     string `json:"host,omitempty" yaml:"host" `
	Port     string `json:"port,omitempty" yaml:"port" `
	Name     string `json:"name"`
	Password string `json:"password"`
}
type LogCof struct {
	LogPath  string `json:"logPath,omitempty" yaml:"logPath"`
	LinkName string `json:"linkName,omitempty" yaml:"linkName"`
}
type Captcha struct {
	Prefix  string        `json:"prefix,omitempty"`
	Expired time.Duration `json:"expired,omitempty"`
}
type PostGreConf struct {
	Host     string
	Port     uint16
	Database string
	User     string // default: OS user name
	Password string
}
