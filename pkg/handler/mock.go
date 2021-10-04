package handler

import (
	"github.com/alphablaster/mocker/pkg/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createMock(c *gin.Context)  {
	var input entity.Mock
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Mock.Create(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllMocksResponse struct {
	Data []entity.Mock `json:"data"`
}

func (h *Handler) getAllMocks(c *gin.Context)  {
	mocks, err := h.services.Mock.GetAll()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllMocksResponse{
		Data: mocks,
	})
}

func (h *Handler) editMock(c *gin.Context)  {

}

func (h *Handler) updateMock(c *gin.Context)  {

}

func (h *Handler) deleteMock(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Mock.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}