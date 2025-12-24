package init

import (
	"5/exam/nacos/config"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/yaml.v2"
)

func InitNacos() {
	//创建客户端
	nacosConf := config.AppConf.Nacos
	clientConfig := constant.ClientConfig{
		NamespaceId:         nacosConf.PublicId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
		Username:            nacosConf.NaCosName,
		Password:            nacosConf.Password,
	}
	//保证端口有一个serverConfigs
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      nacosConf.Host,
			ContextPath: "/nacos",
			Port:        uint64(nacosConf.Port),
			Scheme:      "http",
		},
	}
	//nacos配置
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(fmt.Sprintf("create nacos config client failed, err:%v", err))
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacosConf.DataId,
		Group:  nacosConf.Group})
	if err != nil {
		panic(fmt.Sprintf("get nacos config dataId failed, err:%v", err))
	}
	fmt.Println(string(content))
	err = yaml.Unmarshal([]byte(content), &config.AppConf)
	if err != nil {
		panic(fmt.Sprintf("unmarshal nacos config dataId failed, err:%v", err))
	}

}
