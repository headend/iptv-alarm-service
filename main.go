package main

import (
	loggingpb "github.com/headend/iptv-logging-service/alarm_service"
	"log"
)


func main()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	alarm_service.StartAlarmService()
}

