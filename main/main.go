package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"strings"
	"time"
)

func main() {
	UnknownJson(`{"a":1}`)
	UnknownJson(`[{"a":1},{"b":2}]`)
}

func UnknownJson(data string) {
	if data != `` {
		r := strings.NewReader(data)
		dec := json.NewDecoder(r)
		switch data[0] {
		case 91:
			// "[" 开头的Json
			var param []interface{}
			dec.Decode(&param)
			fmt.Println(param)
		case 123:
			// "{" 开头的Json
			param := make(map[string]interface{})
			dec.Decode(&param)
			fmt.Println(param)
		}
	}
}

func nacos() {
	// 从控制台命名空间管理的"命名空间详情"中拷贝 End Point、命名空间 ID
	var endpoint = "10.12.28.26"
	var namespaceId = "service-hi.yaml"
	// 推荐使用 RAM 用户的 accessKey、secretKey
	var accessKey = "nacos"
	var secretKey = "nacos"

	clientConfig := constant.ClientConfig{
		//
		Endpoint:       endpoint + ":8848",
		NamespaceId:    namespaceId,
		AccessKey:      accessKey,
		SecretKey:      secretKey,
		TimeoutMs:      5 * 1000,
		ListenInterval: 30 * 1000,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	var dataId = "com.alibaba.nacos.example.properties"
	var group = "DEFAULT_GROUP"

	// 发布配置
	success, err := configClient.PublishConfig(vo.ConfigParam{
		DataId:  dataId,
		Group:   group,
		Content: "connectTimeoutInMills=3000"})

	if success {
		fmt.Println("Publish config successfully.")
	}

	time.Sleep(3 * time.Second)

	// 获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})

	fmt.Println("Get config：" + content)

	// 监听配置
	configClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("ListenConfig group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})

	// 删除配置
	success, err = configClient.DeleteConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})

	if success {
		fmt.Println("Delete config successfully.")
	}
}
