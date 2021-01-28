package telegram

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	loggingpb "github.com/headend/iptv-logging-service/proto"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// Update is the type of request that telegram sends once u send message to the bot
type Update struct {
	UpdateID      int           `json:"update_id"`
	Message       Message       `json:"message"`
	CallbackQuery CallbackQuery `json:"callback_query"`
}

// Message is the structure of the message sent to the bot
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
	Date int    `json:"date"`
}

// Chat indicates the conversation to which the message belongs.
type Chat struct {
	ID int `json:"id"`
}

// User is a telegram user
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
}

// CallbackQuery gives the structure of the callback that is received once user clicks on a button
type CallbackQuery struct {
	ID   string `json:"id"`
	From User   `json:"from"`
	Data string `json:"data"`
}

const chatid = -585024223
const telegramAPIBaseURL string = "https://api.telegram.org/bot"
const telegramAPISendMessage string = "/sendMessage"
const telegramTokenEnv string = "1531125523:AAFKXNH6yoL7sr54W8FvCEViOkBFPvocqIM"

// TelegramAPI is the api to which we should send the message to
var TelegramAPI string = telegramAPIBaseURL + telegramTokenEnv + telegramAPISendMessage


// ParseTelegramUpdate takes in the request from telegram and parses Update from it
func ParseTelegramUpdate(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		return nil, err
	}

	return &update, nil
}

// SendTextToTelegram sends text to the user
func SendTextToTelegram(chatID int, text string) (string, error) {
	log.Printf("Sending to chat_id: %d", chatID)
	// 	log.Printf(string(keyboard))
	log.Printf(text)
	log.Println(TelegramAPI)
	response, err := http.PostForm(
		TelegramAPI,
		url.Values{
			"chat_id":      {strconv.Itoa(chatID)},
			"text":         {text},
			// 			"parse_mode":   {"HTML"},
			// 			"reply_markup": {string(keyboard)},
		},
	)

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}

func SendMsgToTelegram(msg string) {
	telegramMsg := ""
	var logData loggingpb.MonitorLogsRequest
	if err := jsonpb.UnmarshalString(msg, &logData); err != nil {
		log.Println(err)
		return
	}
	if logData.Description == "" {
		log.Printf("Log does not content %#v \n", logData)
		return
	}
	switch logData.AfterStatus {
	case 0:
		telegramMsg = fmt.Sprintf("[Critical] %s\n", logData.Description)
	case 1:
		telegramMsg = fmt.Sprintf("[Info] %s\n", logData.Description)
	case 2:
		telegramMsg = fmt.Sprintf("[Critical] %s\n", logData.Description)
	case 3:
		telegramMsg = fmt.Sprintf("[Critical] %s\n", logData.Description)
	default:
		telegramMsg = fmt.Sprintf("[Error] %s\n", logData.Description)
	}
	_, err2 := SendTextToTelegram(chatid, telegramMsg)
	if err2 != nil {
		println(err2)
		return
	}
}
