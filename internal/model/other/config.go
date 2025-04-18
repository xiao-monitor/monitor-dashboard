package other

type Config struct {
	System   System   `yaml:"system"`
	Database Database `yaml:"database"`
}

// Database 数据库配置
type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
	SSLMode  string `yaml:"ssl_mode"`
}

// System 系统配置
type System struct {
	Port int `yaml:"port"`
	Mode string `yaml:"mode"`
}
