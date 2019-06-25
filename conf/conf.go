package conf

type Config struct {
	DbAddr string
}

var Conf = &Config{""}

func init() {
	Conf.DbAddr = "root:123456@tcp(127.0.0.1:3306)/mysql?charset=utf8"
}
