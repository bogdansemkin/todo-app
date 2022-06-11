package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	todo_app "todo-app"
)

func (h *Handler) createList(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input todo_app.TodoList
	if err := c.BindJSON(&input); err != nil {
		fmt.Println("Error is here")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllLists(ctx *gin.Context){

}

func (h *Handler) getListById(ctx *gin.Context){

}

func (h *Handler) updateList(ctx *gin.Context){

}

func (h *Handler) deleteList(ctx *gin.Context){

}