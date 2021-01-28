package main

import (
	"github.com/headend/iptv-alarm-service/alarm-service"
	"log"
)


func main()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	alarm_service.StartAlarmService()
}

