package config

import "github.com/spf13/viper"

type dbConf struct {
	Type         string `mapstructure:"type" json:"type" yaml:"type"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"logMode"`
}

type wxApp struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Appid  string `mapstructure:"appid" json:"appid" yaml:"appid"`
}

type redisCfg struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

func InitDBConfig(cfg *viper.Viper) *dbConf {
	cfgDb := cfg.Sub("database")
	return &dbConf{
		Type:         cfgDb.GetString("type"),
		Username:     cfgDb.GetString("username"),
		Password:     cfgDb.GetString("password"),
		Path:         cfgDb.GetString("path"),
		Dbname:       cfgDb.GetString("dbname"),
		Config:       cfgDb.GetString("config"),
		MaxIdleConns: cfgDb.GetInt(" max-idle-conns"),
		MaxOpenConns: cfgDb.GetInt(" max-open-conns"),
		LogMode:      cfgDb.GetBool("log-mode"),
	}
}

func InitWxAppConfig(cfg *viper.Viper) *wxApp {
	cfgWx := cfg.Sub("wxapp")
	return &wxApp{
		Secret: cfgWx.GetString("secret"),
		Appid:  cfgWx.GetString("appid"),
	}
}

func InitRedisConfig(cfg *viper.Viper) *redisCfg {
	cfgWx := cfg.Sub("redis")
	return &redisCfg{
		Addr:     cfgWx.GetString("addr"),
		Password: cfgWx.GetString("password"),
	}
}
