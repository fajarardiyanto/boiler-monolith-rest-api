package interfaces

type SQLConfig struct {
	Enable     bool   `yaml:"enable" default:"false" desc:"lib:sql:enable"`
	Driver     string `yaml:"driver" default:"mysql" desc:"lib:sql:driver"`
	Host       string `yaml:"host" default:"127.0.0.1" desc:"lib:sql:host"`
	Port       int    `yaml:"port" default:"3306" desc:"lib:sql:port"`
	Username   string `yaml:"username" default:"root"  desc:"lib:sql:username"`
	Password   string `yaml:"password" default:"root" desc:"lib:sql:password"`
	Database   string `yaml:"database" default:"mydb" desc:"lib:sql:database"`
	Options    string `yaml:"options" default:"" desc:"lib:sql:options"`
	Connection string `yaml:"connection" default:"" desc:"lib:sql:connection"`
}
