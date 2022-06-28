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

	// TODO: check if a user has permission to view the list

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

	// checking if the task exists, and getting the task from the db
	task, exists := model.TaskExists(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task with this ID does not exist.",
		})
		return
	}

	// TODO: check if the user has permission to view the list the task is in

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

	// checking if the task exists
	existingTask, exists := model.TaskExists(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task with this ID does not exist.",
		})
		return
	}

	// checking if the user is logged in
	user, err := model.GetLoggedInUser(*c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not logged in.",
		})
		return
	}

	// checking if the user has permission
	if user.Id != existingTask.CreatedById {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to do this.",
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

	// checking if the user is logged in
	user, err := model.GetLoggedInUser(*c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not logged in.",
		})
		return
	}

	// TODO: checking if the list exists
	// TODO: check if the user has permission to create in the list

	// changing the CreatedById
	task.CreatedById = user.Id

	// creating the task
	task, err = model.CreateTask(task)
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

	// checking if the task exists
	existingTask, exists := model.TaskExists(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task with this ID does not exist.",
		})
		return
	}

	// checking if the user is logged in
	user, err := model.GetLoggedInUser(*c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not logged in.",
		})
		return
	}

	// checking if the user has permission
	if user.Id != existingTask.CreatedById {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to do this.",
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

	// changing the task id and the CreatedById
	task.Id = id
	task.CreatedById = user.Id

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

	// checking if the task exists
	existingTask, exists := model.TaskExists(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task with this ID does not exist.",
		})
		return
	}

	// checking if the user is logged in
	user, err := model.GetLoggedInUser(*c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not logged in.",
		})
		return
	}

	// checking if the user has permission
	if user.Id != existingTask.CreatedById {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to do this.",
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
