package controller

import (
	"net/http"
	"strconv"

	"github.com/0l1v3rr/todo/app/model"
	"github.com/0l1v3rr/todo/app/util"
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
// @Failure      401  {object}  util.Error "If the user is not logged in."
// @Failure      403  {object}  util.Error "If the user doesn't have permission to view the list."
// @Failure      404  {object}  util.Error "If the list with this id does not exist."
// @Failure      500  {object}  util.Error "If there was a db error.."
// @Router       /tasks/list/{id} [get]
func GetTasksByListId(c *gin.Context) {
	// parsing the listId parameter
	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.Error{Message: "Please specify a valid id."})
		return
	}

	// checking whether the list exists
	_, exists := model.ListExists(listId)
	if !exists {
		c.JSON(http.StatusNotFound, util.Error{Message: "List with this ID does not exist."})
		return
	}

	// checking if the user is logged in
	user, err := model.GetLoggedInUser(*c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, util.Error{Message: "You are not logged in."})
		return
	}

	// checking if the user has permission to view the list
	listOwner := model.GetListOwnerId(listId)
	if user.Id != listOwner {
		c.JSON(http.StatusForbidden, util.Error{Message: "You do not have the permission to view this list."})
		return
	}

	// getting the tasks from the db
	tasks, err := model.GetTasks(listId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// @Summary      Get tasks by URL
// @Description  Returns the task with the specified url
// @Tags         Task endpoints
// @Produce      json
// @Param 		 url path string true "task URL"
// @Success      200  {object}  model.Task
// @Failure      404  {object}  util.Error "If the task doesn't exist."
// @Failure      401  {object}  util.Error "If the user is not logged in."
// @Failure      403  {object}  util.Error "If the user doesn't have permission to view the list the task is in."
// @Router       /tasks/{url} [get]
func GetTaskByUrl(c *gin.Context) {
	// getting the url from the parameter
	url := c.Param("url")

	task, err := model.GetTaskByUrl(url)
	if err != nil {
		c.JSON(http.StatusNotFound, util.Error{Message: "Task with this URL does not exist."})
		return
	}

	// checking if the user is logged in
	user, err := model.GetLoggedInUser(*c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, util.Error{Message: "You are not logged in."})
		return
	}

	// checking if the user has permission to view the list the task is in
	listOwner := model.GetListOwnerId(task.ListId)
	if user.Id != listOwner {
		c.JSON(http.StatusForbidden, util.Error{Message: "You do not have the permission to view this list."})
		return
	}

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
		c.JSON(http.StatusBadRequest, util.Error{Message: "Please specify a valid id."})
		return
	}

	// checking if the task exists
	existingTask, exists := model.TaskExists(id)
	if !exists {
		c.JSON(http.StatusNotFound, util.Error{Message: "Task with this ID does not exist."})
		return
	}

	// checking if the user is logged in
	user, err := model.GetLoggedInUser(*c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, util.Error{Message: "You are not logged in."})
		return
	}

	// checking if the user has permission
	if user.Id != existingTask.CreatedById {
		c.JSON(http.StatusForbidden, util.Error{Message: "You do not have permission to do this."})
		return
	}

	// changing the IsDone parameter
	task, err := model.ChangeIsDone(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Error{Message: err.Error()})
		return
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
// @Failure      403  {object}  util.Error "If the user doesn't have permission to create task in the list."
// @Failure      404  {object}  util.Error "If the list with the specified ID does not exist."
// @Failure      500  {object}  util.Error "If there was a db error."
// @Router       /tasks/{id} [post]
func CreateTask(c *gin.Context) {
	// binding the task from the body
	var task model.Task

	if err := c.ShouldBindBodyWith(&task, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, util.Error{Message: "Please provide a valid task."})
		return
	}

	// validating the task
	valid, msg := task.Validate()
	if !valid {
		c.JSON(http.StatusBadRequest, util.Error{Message: msg})
		return
	}

	// checking if the user is logged in
	user, err := model.GetLoggedInUser(*c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, util.Error{Message: "You are not logged in."})
		return
	}

	// checking if the list exists
	_, exists := model.ListExists(task.ListId)
	if !exists {
		c.JSON(http.StatusNotFound, util.Error{Message: "List with this ID does not exist."})
		return
	}

	// checking if the user has permission to create task in the list
	listOwner := model.GetListOwnerId(task.ListId)
	if user.Id != listOwner {
		c.JSON(http.StatusForbidden, util.Error{Message: "You do not have the permission to create in this list."})
		return
	}

	// changing the CreatedById
	task.CreatedById = user.Id

	// creating the task
	task, err = model.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Error{Message: err.Error()})
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
		c.JSON(http.StatusBadRequest, util.Error{Message: "Please specify a valid id."})
		return
	}

	// checking if the task exists
	existingTask, exists := model.TaskExists(id)
	if !exists {
		c.JSON(http.StatusNotFound, util.Error{Message: "Task with this ID does not exist."})
		return
	}

	// checking if the user is logged in
	user, err := model.GetLoggedInUser(*c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, util.Error{Message: "You are not logged in."})
		return
	}

	// checking if the user has permission
	if user.Id != existingTask.CreatedById {
		c.JSON(http.StatusForbidden, util.Error{Message: "You do not have permission to do this."})
		return
	}

	// binding the task from the body
	var task model.Task

	if err := c.ShouldBindBodyWith(&task, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, util.Error{Message: "Please provide a valid task."})
		return
	}

	// validating the task
	valid, msg := task.Validate()
	if !valid {
		c.JSON(http.StatusBadRequest, util.Error{Message: msg})
		return
	}

	// changing the task id and the CreatedById
	task.Id = id
	task.CreatedById = user.Id

	// saving the task in the db
	saved, err := model.EditTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Error{Message: err.Error()})
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
		c.JSON(http.StatusBadRequest, util.Error{Message: "Please specify a valid id."})
		return
	}

	// checking if the task exists
	existingTask, exists := model.TaskExists(id)
	if !exists {
		c.JSON(http.StatusNotFound, util.Error{Message: "Task with this ID does not exist."})
		return
	}

	// checking if the user is logged in
	user, err := model.GetLoggedInUser(*c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, util.Error{Message: "You are not logged in."})
		return
	}

	// checking if the user has permission
	if user.Id != existingTask.CreatedById {
		c.JSON(http.StatusForbidden, util.Error{Message: "You do not have permission to do this."})
		return
	}

	// deleting the task
	err = model.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Error{Message: err.Error()})
		return
	}

	// success
	c.Status(http.StatusAccepted)
}
