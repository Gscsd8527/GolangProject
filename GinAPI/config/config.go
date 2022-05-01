package config

// 数据库的结构体
type DataSourceConfig struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	Database   string `mapstructure:"database"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	Charset    string `mapstructure:"charset"`
}

// 整个项目的配置
type ServerConfig struct {
	DataSource DataSourceConfig `mapstructure:"datasource"`
}