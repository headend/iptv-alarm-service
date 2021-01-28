package alarm_service

import (
	"github.com/headend/iptv-alarm-service/telegram"
)

func StartAlarmService()  {
	telegram.SendMsgToTelegram("kaka")
}
