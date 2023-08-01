package config

type Configuration struct {
	Server   Server
	Database Database
	Log      Log
}

type Server struct {
	Port         string `json:"port" yaml:"port" mapstructure:"port"`
	Upload       string `json:"upload" yaml:"upload" mapstructure:"upload"`
	AccessExpire int64  `json:"accessExpire" yaml:"accessExpire" mapstructure:"accessExpire"`
}

type Database struct {
	UserName string `json:"username" yaml:"username" mapstructure:"username"`
	PassWord string `json:"password" yaml:"password" mapstructure:"password"`
	Address  string `json:"address" yaml:"address" mapstructure:"address"`
	Port     string `json:"port" yaml:"port" mapstructure:"port"`
	DBname   string `json:"dbname" yaml:"dbname" mapstructure:"dbname"`
}

type Log struct {
	Root_dir  string `json:"root_dir"`
	FileName  string `json:"filename`
	MaxSize   int    `json:"max_size`
	MaxBackup int    `json:"max_backup`
	MaxAge    int    `json:"max_age"`
	Level     string `json:"level"`
}
