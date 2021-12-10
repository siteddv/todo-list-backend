package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todolistBackend/pkg/model"
)

// @Summary Create item model
// @Security ApiKeyAuth
// @Tags items
// @Description create item model
// @ID create-item
// @Accept  json
// @Produce  json
// @Param input body model.TodoItem true "item info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:list_id/items [post]
func (h *Handler) createItem(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input model.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := h.services.TodoItem.Create(listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": itemId,
	})
}

// @Summary Get All Items
// @Security ApiKeyAuth
// @Tags items
// @Description get all items
// @ID get-all-items
// @Accept  json
// @Produce  json
// @Success 200 {object} GetAllItemsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:list_id/ [get]
func (h *Handler) getAllItems(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetAllItemsResponse{
		Data: items,
	})

}

// @Summary Get Item By Id
// @Security ApiKeyAuth
// @Tags items
// @Description get item by id
// @ID get-item-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} model.TodoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:list_id/items/:item_id/ [get]
func (h *Handler) getItemById(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	item, err := h.services.TodoItem.GetById(listId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Update Item By Id
// @Security ApiKeyAuth
// @Tags items
// @Description update item by id
// @ID put-item-by-id
// @Accept  json
// @Produce  json
// @Param input body model.TodoItem true "item info"
// @Success 200 {string} string 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:list_id/items/:item_id/ [put]
func (h *Handler) updateItem(c *gin.Context) {

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	var input model.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoItem.Update(listId, itemId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete Item By ID
// @Security ApiKeyAuth
// @Tags items
// @Description delete item by id
// @ID delete-item-by-id
// @Accept  json
// @Produce  json
// @Success 200 {string} string 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:list_id/items/:item_id/ [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	err = h.services.TodoItem.DeleteById(listId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
