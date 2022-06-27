package controller

import (
	"net/http"
	"strconv"

	"github.com/0l1v3rr/todo/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetTasksByListId(c *gin.Context) {
	// parsing the listId parameter
	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please specify a valid id.",
		})
		return
	}

	// getting the tasks from the db
	tasks, err := model.GetTasks(listId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	// parsing the id parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please specify a valid id.",
		})
		return
	}

	// getting the task from the db
	task, err := model.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, task)
}

func ChangeTaskStatus(c *gin.Context) {
	// parsing the id parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please specify a valid id.",
		})
		return
	}

	// changing the IsDone parameter
	task, err := model.ChangeIsDone(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	// success
	c.JSON(http.StatusAccepted, task)
}

func CreateTask(c *gin.Context) {
	// binding the task from the body
	var task model.Task

	if err := c.ShouldBindBodyWith(&task, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a valid task.",
		})
		return
	}

	// validating the task
	valid, msg := task.Validate()
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})
		return
	}

	// TODO: checking if the list exists

	// creating the task
	task, err := model.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func EditTask(c *gin.Context) {
	// parsing the id parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please specify a valid id.",
		})
		return
	}

	// binding the task from the body
	var task model.Task

	if err := c.ShouldBindBodyWith(&task, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a valid task.",
		})
		return
	}

	// validating the task
	valid, msg := task.Validate()
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})
		return
	}

	// changing the task id
	task.Id = id

	// saving the task in the db
	saved, err := model.EditTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// success
	c.JSON(http.StatusAccepted, saved)
}

func DeleteTask(c *gin.Context) {
	// parsing the id parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please specify a valid id.",
		})
		return
	}

	// deleting the task
	err = model.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// success
	c.Status(http.StatusAccepted)
}
