package conf

type Config struct {
	DbAddr string
}

var Conf = &Config{""}

func init() {
	Conf.DbAddr = "root:apisapis@tcp(apis.2.chensmallx.top:3306)/mysql?charset=utf8"
}
