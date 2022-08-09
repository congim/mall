package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/zeromicro/go-zero/core/conf"

	"github.com/philchia/agollo/v4"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/congim/mall/apps/app/api/internal/config"
	"github.com/congim/mall/apps/app/api/internal/handler"
	"github.com/congim/mall/apps/app/api/internal/svc"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api-api.yaml", "the config file")

func main() {
	flag.Parse()
	loadConfigFromApollo()
	var c config.Config
	//conf.MustLoad(*configFile, &c)
	conf.LoadFromYamlBytes([]byte(agollo.GetContent(agollo.WithNamespace("chatcheck.yaml"))), &c)
	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func init() {
	logx.DisableStat()
}

func loadConfigFromApollo() error {
	if err := agollo.Start(&agollo.Conf{
		AppID:          "xxx_chatcheck",
		Cluster:        "default",
		NameSpaceNames: []string{"application.properties", "chatcheck.yaml"},
		MetaAddr:       "dev-xxx.xxxx.com:80",
		//AccesskeySecret: "877eae39444e4937b169a59385e3c24f",
	}); err != nil {
		logx.Error("apollo start fail", logx.Field("error", err))
		return err
	}
	if subErr := agollo.SubscribeToNamespaces("chatcheck.yaml"); subErr != nil {
		logx.Error("subscribe name space fail", logx.Field("error", subErr))
		return subErr
	}

	nocitc := make(chan *agollo.ChangeEvent, 100)

	agollo.OnUpdate(func(event *agollo.ChangeEvent) {
		// 监听配置变更
		log.Printf("event:%#v\n", event)
		nocitc <- event
	})

	go func() {
		for {
			select {
			case event, ok := <-nocitc:
				if ok {
					fmt.Println("apollo 配置变更", event.String())
				}
			default:
				fmt.Println("default")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	log.Println("初始化Apollo配置成功")

	// 从默认的application.properties命名空间获取key的值
	val := agollo.GetString("timeout")
	log.Println(val)
	// 获取命名空间下所有key
	keys := agollo.GetAllKeys(agollo.WithNamespace("chatcheck.yaml"))
	fmt.Println(keys)
	// 获取指定一个命令空间下key的值
	other := agollo.GetString("content", agollo.WithNamespace("chatcheck.yaml"))
	log.Println(other)
	// 获取指定命名空间下的所有内容
	namespaceContent := agollo.GetContent(agollo.WithNamespace("chatcheck.yaml"))
	log.Println(namespaceContent)
	return nil
}
