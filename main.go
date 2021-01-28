package main

import (
	"./telegram"
)

const chatid = -585024223

func main()  {
	telegram.SendTextToTelegram(chatid, "Fom bot 2")
}
