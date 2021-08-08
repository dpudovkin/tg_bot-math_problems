package models

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
	User User   `json:"from"`
}
type Chat struct {
	Id int `json:"id"`
}

type User struct {
	Id int64 `json:"id"`
}

type Response struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatId      int                 `json:"chat_id"`
	Text        string              `json:"text"`
	ReplyMarkup ReplyKeyboardMarkup `json:"reply_markup"`
}

type ReplyKeyboardMarkup struct {
	KeyboardButtons [][]KeyboardButton `json:"keyboard"`
}

type KeyboardButton struct {
	Text string `json:"text"`
}
