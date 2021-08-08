package controllers

import (
	models "github.com/p134d/tg_bot-math_problems/models"
)

const (
	NewProblem    = "New problem"
	MyScore       = "MyScore"
	ChangeProblem = "Change problem"
)

func CreateDefaultKeyoard() models.ReplyKeyboardMarkup {
	return models.ReplyKeyboardMarkup{
		KeyboardButtons: [][]models.KeyboardButton{
			{models.KeyboardButton{Text: NewProblem}},
			{models.KeyboardButton{Text: MyScore}},
			//	{models.KeyboardButton{"New problem"}},
		},
	}
}

func CreateProblemKeyBoard() models.ReplyKeyboardMarkup {
	return models.ReplyKeyboardMarkup{
		KeyboardButtons: [][]models.KeyboardButton{
			{models.KeyboardButton{Text: ChangeProblem}},
			{models.KeyboardButton{Text: MyScore}},
			//	{models.KeyboardButton{"New problem"}},
		},
	}
}
