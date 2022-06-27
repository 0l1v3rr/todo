package model

import (
	"fmt"
	"time"

	"github.com/0l1v3rr/todo/app/util"
)

// task struct
type Task struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	ListId      int       `json:"list_id" gorm:"not null"`
	CreatedById int       `json:"created_by_id" gorm:"not null"`
	Title       string    `json:"title" gorm:"not null"`
	Url         string    `json:"url" gorm:"not null"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
}

func (task Task) Validate() (bool, string) {
	// if the title is less than 3 characters
	if len(task.Title) < 3 {
		return false, "The title has to be at least 3 characters long."
	}

	// if the title is too long
	if len(task.Title) > 32 {
		return false, "The title can be maximum 32 characters long."
	}

	// if the description is too long
	if len(task.Description) > 256 {
		return false, "The description can be maximum 256 characters long."
	}

	// if the task is valid
	return true, ""
}

func GetTasks(listId int) ([]Task, error) {
	var tasks []Task

	// getting the tasks from the db where the list id is the specified
	// the result-set should be ordered in descending order by created_at
	tx := DB.Where("list_id = ?", listId).Order("created_at DESC").Find(&tasks)
	if tx.Error != nil {
		return []Task{}, tx.Error
	}

	return tasks, nil
}

func GetTaskById(id int) (Task, error) {
	var task Task

	// getting the task from the db by id
	tx := DB.Where("id = ?", id).First(&task)
	if tx.Error != nil {
		return Task{}, tx.Error
	}

	return task, nil
}

func CreateTask(task Task) (Task, error) {
	// overriding the necessary values
	task.Url = fmt.Sprintf("%s-%d", util.CreateUrlByTitle(task.Title), task.Id)
	task.CreatedAt = time.Now()

	// creating the task in the db
	tx := DB.Create(&task)
	return task, tx.Error
}

func EditTask(task Task) (Task, error) {
	// saving the new task in the db
	tx := DB.Save(&task)
	return task, tx.Error
}

func ChangeIsDone(id int) (Task, error) {
	// getting the task by id
	task, err := GetTaskById(id)
	if err != nil {
		return Task{}, err
	}

	// changing the isDone value to its opposite
	task.IsDone = !task.IsDone

	// saving the task in the db
	task, err = EditTask(task)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func DeleteTask(id int) error {
	// deleting the task from the db
	tx := DB.Unscoped().Delete(&Task{}, id)
	return tx.Error
}
