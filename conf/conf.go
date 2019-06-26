package conf

type Config struct {
	DbAddr string
}

var Conf = &Config{""}

func init() {
	Conf.DbAddr = "apis:apisapis@tcp(apis.2.chensmallx.top:3306)/APIS_DB?charset=utf8&parseTime=true"
}
