package models

import (
	"fmt"
	database "p134d/tg_math_bot/database"

	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	ID          int
	Title       string
	Description string
	Answer      float64
	Cost        int
}

func CreateProblem(problem *Problem) (err error) {
	err = database.Db.Create(problem).Error
	if err != nil {
		return err
	}
	return nil
}

// func GetAllProblems() ([]Problem, error) {
// 	return nil, nil
// }

// func GetProblemById(id string) (Problem, error) {
// 	return Problem{}, nil
// }

func FindProblemById(id int64) Problem {
	problem := Problem{ID: -1}
	database.Db.First(&problem, id)
	return problem
}

func UpdateProblem(Problem *Problem) {
	database.Db.Save(Problem)
}

//TODO
func GetRandomProblem() Problem {
	var problem Problem
	database.Db.First(&problem)
	return problem
}

func (problem *Problem) String() string {
	return fmt.Sprintf("%s\n\n%s\n\nCost:%d", problem.Title, problem.Description, problem.Cost)
}

//TBD
// func DeleteProblemById(id string){
//    database.Db.Where("id = ?", id).Delete()
// }

// func getRandomProblem(db *gorm.DB) (err error){

// }
