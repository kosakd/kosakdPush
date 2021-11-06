package conf

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	fmt.Println(Conf.ServerYaml)
	fmt.Println(Conf.VxPushYaml)
	fmt.Println(Conf.ServerYaml.Ip)
	fmt.Printf("Addr: %s\n", Conf.ServerYaml.Ip+":"+Conf.ServerYaml.Port)
	fmt.Printf("ReadTimeout: %d\n", Conf.ServerYaml.ReadTimeout)
	fmt.Printf("WriteTimeout: %d\n", Conf.ServerYaml.WriteTimeout)
	fmt.Printf("Appid: %s\n", Conf.VxPushYaml.Appid)
	fmt.Printf("Secret: %s\n", Conf.VxPushYaml.Secret)
	fmt.Printf("Templateid: %s\n", Conf.VxPushYaml.Templateid)
	fmt.Printf("Url: %s\n", Conf.VxPushYaml.Url)
}
