package alarm_service

import (
	"github.com/headend/iptv-alarm-service/telegram"
	messagequeue "github.com/headend/share-module/MQ"
	"github.com/headend/share-module/configuration"
	"log"
	"time"
)

func StartAlarmService()  {
	var conf configuration.Conf
	conf.LoadConf()
	var mq messagequeue.MQ
	mq.InitConsumerByTopic(&conf, conf.MQ.MonitorLogsTopic)
	defer mq.CloseConsumer()
	if mq.Err != nil {
		log.Print(mq.Err)
	}
	for  {
		msg, err := mq.Consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Consumer error: %v\n", err)
			log.Print("Retry connect...")
			time.Sleep(10*time.Second)
			continue
		}
		go telegram.SendMsgToTelegram(msg.value)
	}
}
