package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/ecklenke9/SandEWebApp/internal/todo"
	"github.com/gin-gonic/gin"
)

func GetTodoListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, todo.Get())
	return
}

func AddTodoHandler(c *gin.Context) {
	todoItem, statusCode, err := convertHTTPBodyToTodo(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	c.JSON(statusCode, gin.H{"id": todo.Add(todoItem.Message)})
	return
}

func DeleteTodoHandler(c *gin.Context) {
	todoID := c.Param("id")
	if err := todo.Delete(todoID); err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, "")
	return
}

func GetTodoByIdHandler(c *gin.Context) {
	todoID := c.Param("id")
	todo, err := todo.GetByID(todoID)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, todo)
	return
}

func CompleteTodoHandler(c *gin.Context) {
	todoItem, statusCode, err := convertHTTPBodyToTodo(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	todo, err := todo.Complete(todoItem.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, todo)
	return
}

func convertHTTPBodyToTodo(httpBody io.ReadCloser) (todo.Todo, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return todo.Todo{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	return convertJSONBodyToTodo(body)
}

func convertJSONBodyToTodo(jsonBody []byte) (todo.Todo, int, error) {
	var todoItem todo.Todo
	err := json.Unmarshal(jsonBody, &todoItem)
	if err != nil {
		return todo.Todo{}, http.StatusBadRequest, err
	}
	return todoItem, http.StatusOK, nil
}
