package main

import (
	"log"

	json "github.com/json-iterator/go"

	"github.com/yun-mu/MicroServicePractice/common"
	"github.com/yun-mu/MicroServicePractice/config"
	pb "github.com/yun-mu/MicroServicePractice/interface-center/out/log"

	"github.com/micro/go-micro/broker"
)

const service = "log"

var (
	topic string
)

func init() {
	topic = config.GetBrokerTopic(service)
}

func main() {
	srv := common.GetMicroServer(service)

	bk := srv.Server().Options().Broker
	// 这里订阅了 一个 topic, 并提供接口处理
	_, err := bk.Subscribe(topic, subLog)
	if err != nil {
		log.Fatalf("sub error: %v\n", err)
	}

	if err = srv.Run(); err != nil {
		log.Fatalf("srv run error: %v\n", err)
	}
}

func subLog(pub broker.Publication) error {
	var logPB *pb.Log
	if err := json.Unmarshal(pub.Message().Body, &logPB); err != nil {
		return err
	}
	log.Printf("[Log]: user_id: %s,  Msg: %v\n", pub.Message().Header["user_id"], logPB)
	return nil
}
