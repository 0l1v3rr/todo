package controller

import (
	"net/http"
	"strconv"

	"github.com/0l1v3rr/todo/app/model"
	"github.com/0l1v3rr/todo/app/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary      Get lists
// @Description  Returns all the lists the specified user has
// @Tags         List endpoints
// @Produce      json
// @Param 		 userId path int true "user ID"
// @Success      200  {array}   model.List
// @Failure      400  {object}  util.Error "If the id is not valid."
// @Failure      401  {object}  util.Error "If the user is not logged in."
// @Failure      403  {object}  util.Error "If the user doesn't have permission to view the list."
// @Failure      500  {object}  util.Error "If there was a db error."
// @Router       /lists/user/{userId} [get]
func GetListsByUserId(c *gin.Context) {
	// parsing the userId parameter
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.Error{Message: "Please specify a valid id."})
		return
	}

	// checking if the user is logged in
	user, err := model.GetLoggedInUser(*c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, util.Error{Message: "You are not logged in."})
		return
	}

	// checking if the user has permission to view the list
	if user.Id != userId {
		c.JSON(http.StatusForbidden, util.Error{Message: "You do not have permission to view this list."})
		return
	}

	// getting the tasks from the db
	lists, err := model.GetLists(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, lists)
}

// @Summary      Get lists
// @Description  Returns all the lists the specified user has
// @Tags         List endpoints
// @Produce      json
// @Param 		 url path string true "list URL"
// @Success      200  {object}   model.List
// @Failure      401  {object}  util.Error "If the user is not logged in."
// @Failure      403  {object}  util.Error "If the user doesn't have permission to view the list."
// @Failure      404  {object}  util.Error "If the list does not exist."
// @Router       /lists/{url} [get]
func GetListByUrl(c *gin.Context) {
	// getting the url from the parameters
	url := c.Param("url")

	// checking if the user is logged in
	user, err := model.GetLoggedInUser(*c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, util.Error{Message: "You are not logged in."})
		return
	}

	// getting the list from the db
	list, err := model.GetListByUrl(url)
	if err != nil {
		c.JSON(http.StatusNotFound, util.Error{Message: "List with this id does not exist."})
		return
	}

	// checking whether the user has permission to view the list
	if user.Id != list.OwnerId {
		c.JSON(http.StatusForbidden, util.Error{Message: "You do not have permission to view this list."})
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary      Create list
// @Description  Creates a new list
// @Tags         List endpoints
// @Accept       json
// @Produce      json
// @Param 		 list body model.List true "Task to create"
// @Success      201  {object}  model.List
// @Failure      400  {object}  util.Error "If the list is not valid."
// @Failure      401  {object}  util.Error "If the user is not logged in."
// @Failure      500  {object}  util.Error "If there was a db error."
// @Router       /lists [post]
func CreateList(c *gin.Context) {
	// binding the list from the body
	var list model.List

	if err := c.ShouldBindBodyWith(&list, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, util.Error{Message: "Please provide a valid list."})
		return
	}

	// validating the list
	isValid, msg := list.Validate()
	if !isValid {
		c.JSON(http.StatusBadRequest, util.Error{Message: msg})
		return
	}

	// checking if the user is logged in
	user, err := model.GetLoggedInUser(*c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, util.Error{Message: "You are not logged in."})
		return
	}

	// changing the OwnerId
	list.OwnerId = user.Id

	// creating the list
	created, err := model.CreateList(list)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Error{Message: err.Error()})
		return
	}

	// success
	c.JSON(http.StatusCreated, created)
}
