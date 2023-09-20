package todo

import (
	"fmt"

	"github.com/igab-dev/gorm-gin-swagger/pkg/database"
	"gopkg.in/validator.v2"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"max=100"`
	Description string `json:"description" validate:"max=5000"`
	Done        bool   `json:"done"`
}

type TodoCreate struct {
	Title       string `json:"title" validate:"max=100"`
	Description string `json:"description" validate:"max=5000"`
	Done        bool   `json:"done"`
}

func todoValidation(todo Todo) error {
	if err := validator.Validate(todo); err != nil {
		return err
	}
	return nil
}

func todoCreateValidation(todo TodoCreate) error {
	if err := validator.Validate(todo); err != nil {
		return err
	}
	return nil
}

func getAll() (todos []Todo, err error) {

	conn, err := database.OpenConnection()
	if err != nil {
		return
	}

	if err = conn.Find(&todos).Error; err != nil {
		return
	}

	return
}

func delete(id int) (err error) {

	conn, err := database.OpenConnection()
	if err != nil {
		return
	}

	var todo Todo
	todo.ID = id

	db := conn.Delete(&todo)
	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected < 1 {
		err = fmt.Errorf("todo doesn't exist, id: %d", id)
		return
	}

	return
}

func byId(id int) (todo Todo, err error) {

	conn, err := database.OpenConnection()

	if err != nil {
		return
	}

	if err = conn.First(&todo, id).Error; err != nil {
		return
	}

	return
}

func new(todo TodoCreate) (Todo, error) {

	var ntodo Todo
	ntodo.Description = todo.Description
	ntodo.Title = todo.Title
	ntodo.Done = todo.Done

	conn, err := database.OpenConnection()
	if err != nil {
		return ntodo, err
	}

	res := conn.Create(&ntodo)
	if err = res.Error; err != nil {
		return ntodo, err
	}

	return ntodo, err
}

func edit(todo Todo) (Todo, error) {

	conn, err := database.OpenConnection()
	if err != nil {
		return todo, err
	}

	if err = conn.Save(&todo).Error; err != nil {
		return todo, err
	}

	return todo, err
}
