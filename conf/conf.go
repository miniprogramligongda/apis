package conf

type Config struct {
	DbAddr    string
	AppID     string
	Secret    string
	CertFile  string
	KeyFile   string
	RedisAddr string
}

// Conf contains the basis infomation of project
var Conf = &Config{}

func init() {
	Conf.DbAddr = "apis:apisapis@tcp(apis.2.chensmallx.top:3306)/APIS_DB?charset=utf8&parseTime=true"
	Conf.AppID = "wx7f5430dfe80d0c10"
	Conf.Secret = "1f49e3c660e1fad1bfad26ed016656d8"
	Conf.RedisAddr = "apis.2.chensmallx.top:6379"
}
