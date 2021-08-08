package main

import (
	server "github.com/p134d/tg_bot-math_problems/controllers/server"
	"github.com/p134d/tg_bot-math_problems/database"
	models "github.com/p134d/tg_bot-math_problems/models"
)

func main() {
	database.InitDb()

	launch()

}

func launch() {

	botToken := "1926510391:AAEPwk-QeXSS2cbZNkK6noMAt0B7WL2_8eY"
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken
	server.Start(botUrl)
}

func initMigration() {
	database.InitDb()
	database.Db.AutoMigrate(&models.UserProblem{}, &models.Problem{})
}
