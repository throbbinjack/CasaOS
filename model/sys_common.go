package model

import "time"

//系统配置
type SysInfoModel struct {
	Name string //系统名称
}

//用户相关
type UserModel struct {
	UserName    string
	PWD         string
	Token       string
	Head        string
	Email       string
	Description string
}

//服务配置
type ServerModel struct {
	HttpPort  string
	RunMode   string
	ServerApi string
}

//服务配置
type APPModel struct {
	LogSavePath    string
	LogSaveName    string
	LogFileExt     string
	DateStrFormat  string
	DateTimeFormat string
	TimeFormat     string
	DateFormat     string
	ProjectPath    string
}

//公共返回模型
type Result struct {
	Success int         `json:"success" example:"200"`
	Message string      `json:"message" example:"ok"`
	Data    interface{} `json:"data" example:"返回结果"`
}

//zeritier相关
type ZeroTierModel struct {
	UserName string
	PWD      string
	Token    string
}

//redis配置文件
type RedisModel struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type SystemConfig struct {
	SearchSwitch    bool   `json:"search_switch"` //搜索开关
	SearchEngine    string `json:"search_engine"` //搜索引擎
	ShortcutsSwitch bool   `json:"shortcuts_switch"`
	WidgetsSwitch   bool   `json:"widgets_switch"`
	BackgroundType  string `json:"background_type"`
	Background      string `json:"background"`
	AutoUpdate      bool   `json:"auto_update"`
}