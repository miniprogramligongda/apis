package conf

type Config struct {
	DbAddr string
	AppID  string
	Secret string
}

// Conf contains the basis infomation of project
var Conf = &Config{"", "", ""}

func init() {
	Conf.DbAddr = "apis:apisapis@tcp(apis.2.chensmallx.top:3306)/APIS_DB?charset=utf8&parseTime=true"
	Conf.AppID = "wx67763b6060ba9528"
	Conf.Secret = ""
}
