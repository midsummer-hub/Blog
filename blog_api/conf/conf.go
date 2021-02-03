package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type conf struct {
	Server server 		`yaml:"server"`
	Db db 				`yaml:"db"`
	MyLog myLog 		`yaml:"myLog"`
	Cache cache 		`yaml:"cache"`

}

type server struct {
	Address string 		`yaml:"address"`
	Model string 		`yaml:"model"`
}

type db struct {
	Dialects string 	`yaml:"dialects"`
	Host string 		`yaml:"host"`
	Port int 			`yaml:"port"`
	Db string 			`yaml:"db"`
	Username string 	`yaml:"username"`
	Password string 	`yaml:"password"`
	Charset string 		`yaml:"charset"`
	MaxIdle int 		`yaml:"maxIdle"`
	MaxOpen int 		`yaml:"maxOpen"`
}


type myLog struct {
	Path string 		`yaml:"path"`
	Name string 		`yaml:"name"`
	Model string 		`yaml:"model"`
}

type cache struct {
	Expire int 			`yaml:"expire"`
	Clearup int			`yaml:"clearup"`
}

var Conf *conf
//配置初始化
func init()  {
	yamkFile,err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamkFile,&Conf)
	if err != nil {
		panic(err)
	}
}



