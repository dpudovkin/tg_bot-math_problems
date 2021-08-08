package models

import (
	database "github.com/p134d/tg_bot-math_problems/database"

	"gorm.io/gorm"
)

type UserProblem struct {
	gorm.Model
	ID              int64 // unique id from telegram message
	Score           int   // sum of score for solved problem
	Current_problem int   //default -1
}

func CreateUser(UserProblem *UserProblem) (err error) {
	err = database.Db.Create(UserProblem).Error
	if err != nil {
		return err
	}
	return nil
}

func AddUser(id int64) {
	if FindUserById(id).ID < 0 {
		user := UserProblem{ID: id, Current_problem: -1}
		CreateUser(&user)
	}
}

func FindUserById(id int64) UserProblem {
	user := UserProblem{}
	database.Db.First(&user, id)
	return user
}
