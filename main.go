package main

import (
	"github.com/headend/iptv-alarm-service/telegram"
)

const chatid = -585024223

func main()  {
	telegram.SendTextToTelegram(chatid, "Fom bot 2")
}
