package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	controllers "github.com/p134d/tg_bot-math_problems/controllers"
	models "github.com/p134d/tg_bot-math_problems/models"
)

func respond(botUrl string, update models.Update) error {
	respondFunction := parseCommand(update.Message.Text)

	preResond(update)
	var botMessage models.BotMessage = respondFunction(update)
	//botMessage.ChatId = update.Message.Chat.Id
	//botMessage.Text = update.Message.Text
	byteArr, err := json.Marshal(botMessage)

	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json",
		bytes.NewReader(byteArr))
	if err != nil {
		return err
	}
	return nil
}

func preResond(update models.Update) {
	models.AddUser(update.Message.User.Id)
}

// get score
func respondScore(update models.Update) models.BotMessage {
	botMessage := models.BotMessage{}
	user := models.FindUserById(int64(update.Message.User.Id))
	botMessage.Text = strconv.FormatInt(int64(user.Score), 10)
	botMessage.ChatId = update.Message.Chat.Id
	botMessage.ReplyMarkup = controllers.CreateDefaultKeyoard() // TODO
	return botMessage
}

// get problem
func respondProblem(update models.Update) models.BotMessage {
	botMessage := models.BotMessage{}
	//user:= models.FindUserById(int64(update.Message.User.Id))
	// TODO return random unsoldev problem for this user
	problem := models.GetRandomProblem()
	botMessage.Text = string(problem.String())
	controllers.AppointProblem(models.FindUserById(update.Message.User.Id), problem)
	botMessage.ChatId = update.Message.Chat.Id
	botMessage.ReplyMarkup = controllers.CreateProblemKeyBoard() // TODO
	return botMessage
}

//get status of answer
func respondAnswer(update models.Update) models.BotMessage {
	botMessage := models.BotMessage{}
	botMessage.ReplyMarkup = controllers.CreateProblemKeyBoard()
	user := models.FindUserById(int64(update.Message.User.Id))
	if user.Current_problem < 0 {
		botMessage.Text = "Unknown command("
	} else {
		problem := models.FindProblemById(int64(user.Current_problem))
		answer, err := strconv.ParseFloat(update.Message.Text, 64)
		if err != nil {
			botMessage.Text = "Incorrect answer("
		} else {
			if problem.Answer == answer {
				botMessage.Text = "Correct answer!"
				controllers.SolveProblem(user, problem)
				botMessage.ReplyMarkup = controllers.CreateDefaultKeyoard()
			}
		}
	}
	botMessage.ChatId = update.Message.Chat.Id
	return botMessage
}

func parseCommand(text string) func(models.Update) models.BotMessage {
	switch text {
	case controllers.MyScore:
		return respondScore
	case controllers.ChangeProblem:
		return respondProblem
	case controllers.NewProblem:
		return respondProblem
	default:
		return respondAnswer
	}
}
