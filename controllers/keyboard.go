package controllers

import (
	models "p134d/tg_math_bot/models"
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
