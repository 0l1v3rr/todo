package controller

import (
	"net/http"
	"strconv"

	"github.com/0l1v3rr/todo/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary      Get list tasks
// @Description  Returns all the tasks in the specified list
// @Tags         Task endpoints
// @Produce      json
// @Param 		 id path int true "list ID"
// @Success      200  {array}   model.Task
// @Failure      400  {object}  util.Error "If the id is not valid."
// @Failure      500  {object}  util.Error "If there was a db error.."
// @Router       /tasks/list/{id} [get]
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

// @Summary      Get tasks
// @Description  Returns the task with the specified id
// @Tags         Task endpoints
// @Produce      json
// @Param 		 id path int true "task ID"
// @Success      200  {object}  model.Task
// @Failure      400  {object}  util.Error "If the id is not valid."
// @Failure      404  {object}  util.Error "If the task doesn't exist."
// @Router       /tasks/{id} [get]
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

// @Summary      Change task status
// @Description  Changes the task status to its opposite value
// @Tags         Task endpoints
// @Produce      json
// @Param 		 id path int true "task ID"
// @Success      202  {object}  model.Task
// @Failure      400  {object}  util.Error "If the id is not valid."
// @Failure      401  {object}  util.Error "If the user is not logged in."
// @Failure      403  {object}  util.Error "If the user has no permission to do this."
// @Failure      404  {object}  util.Error "If the task doesn't exist."
// @Router       /tasks/{id} [patch]
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

// @Summary      Create task
// @Description  Creates a new task in the db
// @Tags         Task endpoints
// @Accept       json
// @Produce      json
// @Param 		 task body model.Task true "Task to create"
// @Success      201  {object}  model.Task
// @Failure      400  {object}  util.Error "If the task is not valid."
// @Failure      401  {object}  util.Error "If the user is not logged in."
// @Failure      500  {object}  util.Error "If there was a db error."
// @Router       /tasks/{id} [post]
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

// @Summary      Edit task
// @Description  Edits the task
// @Tags         Task endpoints
// @Accept       json
// @Produce      json
// @Param 		 task body model.Task true "Task to create"
// @Param 		 id path int true "task ID"
// @Success      202  {object}  model.Task
// @Failure      400  {object}  util.Error "If the task or the id is not valid."
// @Failure      404  {object}  util.Error "If the task does not exist."
// @Failure      401  {object}  util.Error "If the user is not logged in."
// @Failure      403  {object}  util.Error "If the user has no permission to do this."
// @Failure      500  {object}  util.Error "If there was a db error."
// @Router       /tasks/{id} [put]
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

// @Summary      Delete task
// @Description  Deletes the task
// @Tags         Task endpoints
// @Param 		 id path int true "task ID"
// @Success      202  {object}  model.Task
// @Failure      400  {object}  util.Error "If the id is not valid."
// @Failure      404  {object}  util.Error "If the task does not exist."
// @Failure      401  {object}  util.Error "If the user is not logged in."
// @Failure      403  {object}  util.Error "If the user has no permission to do this."
// @Failure      500  {object}  util.Error "If there was a db error."
// @Router       /tasks/{id} [delete]
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
